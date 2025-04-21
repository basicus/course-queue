### День 3 
***Использется  кластер из День 2***

```shell
cd ./kafka-kraft-cluster
docker compose up -d
```

### Цель
- Run Streams on top of the cluster from previous homework
- Create simple mapper for any of the simulated topic
- Implement simple aggregate processor
- Implement join processor

### Подготовка: Сгенерировать данные о Клиентах
Данные будут использоваться для обогащения потока платежей.
```text
python3 hw3/init_customer.py localhost:29092 customers
```

### Запустить producer из hw2 
```shell
python3 producer.py --bootstrap_servers localhost:29092 localhost:29093
```

Данный продюсер будет генерировать поток событий о платежах клиентов.

### Запустить KafkaStreamsAp
Данный класс реализует поток агрегации - среднее
Для этого необходимо исправить/проверить в build.gradle mainClass на course.queues.KafkaStreamsAp

Пример результата работы:
```text
Out << key: [6012@1745187180000/1745187195000]:	11.58 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [3605@1745187180000/1745187195000]:	48.795 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [7240@1745187180000/1745187195000]:	19.915 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [3513@1745187195000/1745187210000]:	45.505 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [9568@1745187195000/1745187210000]:	35.685 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [902@1745187225000/1745187240000]:	24.285 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
[kafka-streams-aggregate-52bcaba2-70ac-4e6f-afad-f4ac5fb7367b-StreamThread-1] INFO org.apache.kafka.streams.state.internals.RocksDBTimestampedStore - Opening store KSTREAM-AGGREGATE-STATE-STORE-0000000003.1745187240000 in regular mode
Out << key: [3513@1745187240000/1745187255000]:	26.27 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [3352@1745187255000/1745187270000]:	27.015 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
[kafka-streams-aggregate-52bcaba2-70ac-4e6f-afad-f4ac5fb7367b-StreamThread-1] INFO org.apache.kafka.streams.state.internals.RocksDBTimestampedStore - Opening store KSTREAM-AGGREGATE-STATE-STORE-0000000003.1745187180000 in regular mode
Out << key: [4874@1745187180000/1745187195000]:	49.935 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [7604@1745187210000/1745187225000]:	38.33 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [6348@1745187210000/1745187225000]:	35.22 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [5917@1745187210000/1745187225000]:	10.18 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [8746@1745187225000/1745187240000]:	21.16 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [419@1745187225000/1745187240000]:	5.43 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
[kafka-streams-aggregate-52bcaba2-70ac-4e6f-afad-f4ac5fb7367b-StreamThread-1] INFO org.apache.kafka.streams.state.internals.RocksDBTimestampedStore - Opening store KSTREAM-AGGREGATE-STATE-STORE-0000000003.1745187240000 in regular mode
Out << key: [4885@1745187240000/1745187255000]:	46.24 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [3754@1745187240000/1745187255000]:	39.585 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [6946@1745187240000/1745187255000]:	45.76 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [4241@1745187255000/1745187270000]:	49.89 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
[kafka-streams-aggregate-52bcaba2-70ac-4e6f-afad-f4ac5fb7367b-StreamThread-1] INFO org.apache.kafka.streams.state.internals.RocksDBTimestampedStore - Opening store KSTREAM-AGGREGATE-STATE-STORE-0000000003.1745187180000 in regular mode
Out << key: [6217@1745187195000/1745187210000]:	5.475 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [9059@1745187195000/1745187210000]:	46.895 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [1463@1745187210000/1745187225000]:	42.35 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [4360@1745187225000/1745187240000]:	20.18 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [4166@1745187225000/1745187240000]:	19.275 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
[kafka-streams-aggregate-52bcaba2-70ac-4e6f-afad-f4ac5fb7367b-StreamThread-1] INFO org.apache.kafka.streams.state.internals.RocksDBTimestampedStore - Opening store KSTREAM-AGGREGATE-STATE-STORE-0000000003.1745187240000 in regular mode
Out << key: [6395@1745187240000/1745187255000]:	14.05 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
Out << key: [7803@1745187255000/1745187270000]:	42.085 (org.apache.kafka.streams.kstream.Windowed, java.lang.Double)
```


### Запустить KafkaStreamsJoinApp
Данный класс реализует поток слияния (join) данных из потока, который генерируется producer.by с данными из GlobalTable

Для этого необходимо исправить в build.gradle mainClass на course.queues.KafkaStreamsJoinApp

Пример результата работы:
```text
In  >> key: 2483:	{"payment_id": 474, "timestamp": 1745186558.5792928, "amount": 83.15, "status": "successful", "customer_id": 2483}
Out << key: 2483:	{"payment_id":474,"timestamp":"1745186558.5792928","amount":83.15,"status":"successful","customer_id":2483,"customer_name":"Amanda Terry"}
In  >> key: 2897:	{"payment_id": 273, "timestamp": 1745186516.5250165, "amount": 95.34, "status": "successful", "customer_id": 2897}
Out << key: 2897:	{"payment_id":273,"timestamp":"1745186516.5250165","amount":95.34,"status":"successful","customer_id":2897,"customer_name":"Anthony Hogan"}
In  >> key: 5935:	{"payment_id": 782, "timestamp": 1745186531.5435617, "amount": 99.43, "status": "failed", "customer_id": 5935}
Out << key: 5935:	{"payment_id":782,"timestamp":"1745186531.5435617","amount":99.43,"status":"failed","customer_id":5935,"customer_name":"Phillip Moody"}
In  >> key: 2912:	{"payment_id": 231, "timestamp": 1745186534.548511, "amount": 79.28, "status": "successful", "customer_id": 2912}
Out << key: 2912:	{"payment_id":231,"timestamp":"1745186534.548511","amount":79.28,"status":"successful","customer_id":2912,"customer_name":"Michael Berg III"}
In  >> key: 7320:	{"payment_id": 555, "timestamp": 1745186543.5570965, "amount": 93.52, "status": "failed", "customer_id": 7320}
Out << key: 7320:	{"payment_id":555,"timestamp":"1745186543.5570965","amount":93.52,"status":"failed","customer_id":7320,"customer_name":"Jason Levine"}
In  >> key: 2739:	{"payment_id": 748, "timestamp": 1745186545.5615878, "amount": 50.42, "status": "failed", "customer_id": 2739}
```

Данные поля customer_name берутся из GlobalTable (customers).  
