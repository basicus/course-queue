package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/spf13/pflag"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kslog"
	"log/slog"
	"os"
	"strings"
)

var UniqueViolationErr = pq.ErrorCode("23505")

func main() {
	var dbPath, seeds, topics, groups string
	pflag.StringVarP(&seeds, "seeds", "s", "localhost:29092,localhost:29093", "Seed for brokers")
	pflag.StringVarP(&topics, "topics", "t", "warehouse.orders.transactions", "Comma separated list of topics")
	pflag.StringVarP(&dbPath, "dsn", "", "postgres://postgres:postgres@localhost/goods_db?sslmode=disable", "Connect string to database: postgres://postgres:postgres@localhost/remains?sslmode=disable")
	pflag.StringVarP(&groups, "groups", "g", "go-group3", "Consumer group")

	pflag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Starting Kafka Brokers: ", seeds)
	seedsOpt := strings.Split(seeds, ",")
	topicsOpt := strings.Split(topics, ",")

	opts := []kgo.Opt{
		kgo.SeedBrokers(seedsOpt...),
		kgo.DisableAutoCommit(),
		kgo.MaxConcurrentFetches(3),
		kgo.WithLogger(kslog.New(logger)),
		kgo.ConsumerGroup(groups),
		kgo.ConsumeTopics(topicsOpt...),
	}
	slog.Info("Kafka broker list %+v", seeds)

	cl, err := kgo.NewClient(opts...)
	if err != nil {
		logger.Error("Unable to create client %s", err)
		panic(err)
	}

	// Database connection
	if dbPath == "" {
		logger.Error("Not defined db connection")
	}
	db, err := sqlx.Open("postgres", dbPath)
	err = db.Ping()
	if err != nil {
		logger.Error("Unable to connect to db %s", err)
		panic(err)
	}

	// Read connections
	for {
		fs := cl.PollFetches(context.Background())
		if fs.IsClientClosed() {
			return
		}
		fs.EachRecord(func(r *kgo.Record) {
			var key = new(Key)

			err = json.Unmarshal(r.Key, &key)
			if err != nil {
				logger.Error("Error unmarshalling key %s", err)
			} else {
				var value = new(Transactions)
				err = json.Unmarshal(r.Value, &value)
				if err != nil {
					logger.Error("Error unmarshalling value %s", err)
				} else {
					/* check if transaction exists */
					tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
					if err != nil {
						logger.Error("Error beginning transaction %s", err)
					}
					exec, err := tx.Exec("INSERT INTO remains.transactions_log (transaction_id) VALUES ($1)", key.Id)

					var pgErr *pq.Error
					if err != nil {
						switch {
						// Using errors.As we inspect the error type
						case errors.As(err, &pgErr):
							// I chose to be explicit and call .Name() rather than
							// case over the Code which in this case is 25302 or something
							switch pgErr.Code.Name() {
							case "unique_violation":
								logger.Info("Skip! Already seen transaction")
							default:
								logger.Error("Error inserting transaction_log %s", err)
								panic("Exiting")
							}
							return
						}
					} else {
						rowsAffected, _ := exec.RowsAffected()
						logger.Info(fmt.Sprintf("Inserted %d tx records", rowsAffected))
						if rowsAffected > 0 {
							var pcs int
							switch value.After.Operation {
							case "add":
								pcs = int(value.After.Pcs)
							case "remove":
								pcs = -1 * int(value.After.Pcs)
							}
							exec, err := tx.Exec("INSERT INTO remains.remains (good_id,cnt) VALUES ($1,$2) ON CONFLICT (good_id) DO UPDATE SET cnt = remains.remains.cnt + $3", value.After.GoodId, pcs, pcs)
							if err != nil {
								logger.Error("Error unmarshalling value %s", err)
								panic("Exiting")
							}
							rowsAffected, _ := exec.RowsAffected()
							logger.Info(fmt.Sprintf("Inserted/updated %d records", rowsAffected))
						}
					}
					err = tx.Commit()
				}
			}
		})
		if err := cl.CommitUncommittedOffsets(context.Background()); err != nil {
			logger.Error("commit records failed: %v", err)
			continue
		}
	}

}
