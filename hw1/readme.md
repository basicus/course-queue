### День 1

#### Задание
Implement simple processing pipeline:
- Run/use any broker of your choice
- Create minimal producer (infinite generation of messages)
- Create minimal consumer (forever consume messages)
- Run several producers and several consumer simultaneously

Для запуска  кластера nats нейобходимо перейти в подпапку **_nats-cluster_** и выполнить команду
```shell
cd ./nats-cluster
docker compose up -d
```

### Запуск консьюмера

В первой консоли выполнить запуск консьюмера

**Внимание!** Должен быть установлен go
```shell
go run consumer/consumer.go
```


### Запуск продьюсера

Во второй консоли выполнить запуск продьюсера
```shell
go run producer/producer.go

```
Пример результата работы продюсера:
```text
2025/04/14 23:57:06 Producer connected to NATS
2025/04/14 23:57:07 Published message: bcdefghijklmnopqrstu
2025/04/14 23:57:08 Published message: vwxyzABCDEFGHIJKLMNO
2025/04/14 23:57:09 Published message: PQRSTUVWXYZ012345678
2025/04/14 23:57:10 Published message: 9bcdefghijklmnopqrst
2025/04/14 23:57:11 Published message: uvwxyzABCDEFGHIJKLMN
2025/04/14 23:57:12 Published message: OPQRSTUVWXYZ01234567
2025/04/14 23:57:13 Published message: 89bcdefghijklmnopqrs
2025/04/14 23:57:14 Published message: tuvwxyzABCDEFGHIJKLM
2025/04/14 23:57:15 Published message: NOPQRSTUVWXYZ0123456
2025/04/14 23:57:16 Published message: 789bcdefghijklmnopqr
2025/04/14 23:57:17 Published message: stuvwxyzABCDEFGHIJKL
2025/04/14 23:57:18 Published message: MNOPQRSTUVWXYZ012345
2025/04/14 23:57:19 Published message: 6789bcdefghijklmnopq
2025/04/14 23:57:20 Published message: rstuvwxyzABCDEFGHIJK
2025/04/14 23:57:21 Published message: LMNOPQRSTUVWXYZ01234
2025/04/14 23:57:22 Published message: 56789bcdefghijklmnop
2025/04/14 23:57:23 Published message: qrstuvwxyzABCDEFGHIJ
2025/04/14 23:57:24 Published message: KLMNOPQRSTUVWXYZ0123
```

Пример результата работы консьюмера:
```text
go run consumer/consumer.go
2025/04/14 23:57:53 Consumer connected to NATS
Received message: {Timestamp:2025-04-14 23:57:54.238043077 +0300 MSK MessageId:4be07ad2-69e7-4b4b-b87a-4b0db6084c15 MessageText:ABCDEFGHIJKLMNOPQRST}
Received message: {Timestamp:2025-04-14 23:57:55.237621649 +0300 MSK MessageId:11f9dbc7-eafe-4f9b-aa12-db9b0d278ae4 MessageText:UVWXYZ0123456789bcde}
Received message: {Timestamp:2025-04-14 23:57:56.238016597 +0300 MSK MessageId:3d395d61-c040-40a0-ae9f-d47e9f1b301f MessageText:fghijklmnopqrstuvwxy}
Received message: {Timestamp:2025-04-14 23:57:57.237503021 +0300 MSK MessageId:3050b445-e679-4211-b177-ddd4b7d534f1 MessageText:zABCDEFGHIJKLMNOPQRS}
Received message: {Timestamp:2025-04-14 23:57:58.23739215 +0300 MSK MessageId:1876809b-532e-4745-823d-56db86d1d37b MessageText:TUVWXYZ0123456789bcd}
Received message: {Timestamp:2025-04-14 23:57:59.237977052 +0300 MSK MessageId:191b6470-7b27-470f-9dfd-8a4401773584 MessageText:efghijklmnopqrstuvwx}
Received message: {Timestamp:2025-04-14 23:58:00.237276791 +0300 MSK MessageId:d21363a9-0dd8-4706-a898-eb1c1888c972 MessageText:yzABCDEFGHIJKLMNOPQR}
Received message: {Timestamp:2025-04-14 23:58:01.237723508 +0300 MSK MessageId:9685c322-7876-45f7-9e85-22c9368291a8 MessageText:STUVWXYZ0123456789bc}
Received message: {Timestamp:2025-04-14 23:58:02.237248402 +0300 MSK MessageId:99bacc0c-7f3c-4de7-914f-b38d499a2d02 MessageText:defghijklmnopqrstuvw}
Received message: {Timestamp:2025-04-14 23:58:03.237479572 +0300 MSK MessageId:6481887e-6b0e-418c-8e45-03afaff188c4 MessageText:xyzABCDEFGHIJKLMNOPQ}
Received message: {Timestamp:2025-04-14 23:58:04.23727165 +0300 MSK MessageId:6a0a5c8c-6ad1-46cc-9e53-c01eb5dc39e1 MessageText:RSTUVWXYZ0123456789b}
Received message: {Timestamp:2025-04-14 23:58:05.23760462 +0300 MSK MessageId:f6ef2ff5-8866-4ac4-a9a0-90948303b34c MessageText:cdefghijklmnopqrstuv}
Received message: {Timestamp:2025-04-14 23:58:06.237807626 +0300 MSK MessageId:06633748-0d81-48a1-9426-a117ef29752b MessageText:wxyzABCDEFGHIJKLMNOP}
Received message: {Timestamp:2025-04-14 23:58:07.237319852 +0300 MSK MessageId:33106ce7-d376-4c97-9c8f-f91a3c34d86d MessageText:QRSTUVWXYZ0123456789}
Received message: {Timestamp:2025-04-14 23:58:08.237670188 +0300 MSK MessageId:5f7bbbc2-a44d-4b31-bdee-5272304247c3 MessageText:bcdefghijklmnopqrstu}
Received message: {Timestamp:2025-04-14 23:58:09.237308937 +0300 MSK MessageId:15d0f0dd-f30b-49c7-930e-bdbcda421396 MessageText:vwxyzABCDEFGHIJKLMNO}
```
