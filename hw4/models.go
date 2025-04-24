package main

import "time"

type Transaction struct {
	Id        string    `json:"id"`
	GoodId    int       `json:"good_id"`
	Operation string    `json:"operation"`
	Pcs       uint32    `json:"pcs"`
	CreatedAt time.Time `json:"created_at"`
}

type Key struct {
	Id string `json:"id"`
}

type Transactions struct {
	Before Transaction `json:"before"`
	After  Transaction `json:"after"`
	Source struct {
		Version   string      `json:"version"`
		Connector string      `json:"connector"`
		Name      string      `json:"name"`
		TsMs      int64       `json:"ts_ms"`
		Snapshot  string      `json:"snapshot"`
		Db        string      `json:"db"`
		Sequence  string      `json:"sequence"`
		TsUs      int64       `json:"ts_us"`
		TsNs      int64       `json:"ts_ns"`
		Schema    string      `json:"schema"`
		Table     string      `json:"table"`
		TxId      int         `json:"txId"`
		Lsn       int         `json:"lsn"`
		Xmin      interface{} `json:"xmin"`
	} `json:"source"`
	Transaction interface{} `json:"transaction"`
	Op          string      `json:"op"`
	TsMs        int64       `json:"ts_ms"`
	TsUs        int64       `json:"ts_us"`
	TsNs        int64       `json:"ts_ns"`
}
