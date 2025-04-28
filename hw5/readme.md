### День 5

#### Задание
Play with NATS
- Test Pub/Sub
- Test Request/Response
- Test Work Queue (group consumers)

#### Настройка окружения

Необходимо выполнить

```shell
docker compose up -d
```


#### 1. Pub/Sub
Задание по Pub/Sub выполнено в hw1

#### 2. Request/Response
Запуск producer в регионе NN
```shell
go run producer/producer.go -url nats://localhost:4222
```

Запуск consumer в регионе MSK

```shell
go run consumer/consumer.go -url nats://localhost:4233
```

Результат работы:
```text
2025/04/28 07:51:47 Producer connected to NATS  YAR
2025/04/28 07:51:48 >>>[YAR] Published message: {Timestamp:2025-04-28 07:51:48.811741415 +0300 MSK m=+1.002058324 MessageId:a442b08a-1efe-4d46-ac5c-1d2c773b08bb MessageText:bcdefghijklmnopqrstu}
2025/04/28 07:51:48 <<<Reply message: {MessageId:a442b08a-1efe-4d46-ac5c-1d2c773b08bb MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:51:49 >>>[YAR] Published message: {Timestamp:2025-04-28 07:51:49.81197475 +0300 MSK m=+2.002291601 MessageId:64ff9a53-e2ab-4f1b-a85e-c4927afc2f51 MessageText:vwxyzABCDEFGHIJKLMNO}
2025/04/28 07:51:49 <<<Reply message: {MessageId:64ff9a53-e2ab-4f1b-a85e-c4927afc2f51 MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:51:50 >>>[YAR] Published message: {Timestamp:2025-04-28 07:51:50.812600495 +0300 MSK m=+3.002917382 MessageId:24a76ac5-e3b5-4734-aab8-384b8f62d55a MessageText:PQRSTUVWXYZ012345678}
2025/04/28 07:51:50 <<<Reply message: {MessageId:24a76ac5-e3b5-4734-aab8-384b8f62d55a MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:51:51 >>>[YAR] Published message: {Timestamp:2025-04-28 07:51:51.811695654 +0300 MSK m=+4.002012500 MessageId:2120cb5c-1bf3-474e-ada2-b4665b93d404 MessageText:9bcdefghijklmnopqrst}
2025/04/28 07:51:51 <<<Reply message: {MessageId:2120cb5c-1bf3-474e-ada2-b4665b93d404 MessageStatus:200 MessageReply:OK from NN}
```
Необходимо обратить внимание, что producer подключается к YAR кластеру, в то время как consumer подключается к NN кластеру.

Если запустить второй consumer в кластере YAR, то обработка будет выполняться на локальном относительно подключения продюсере кластере
Пример лога после подключения consumer из кластера YAR. В результате запросы стали обрабатываться "локально" в YAR. 
```text
2025/04/28 07:55:40 <<<Reply message: {MessageId:824646b2-9431-4532-b3b0-3119cab9b639 MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:41 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:41.5024а96066 +0300 MSK m=+35.002687497 MessageId:69ed76bc-2c76-4e2d-a38f-ab1537f889cf MessageText:klmnopqrstuvwxyzABCD}
2025/04/28 07:55:41 <<<Reply message: {MessageId:69ed76bc-2c76-4e2d-a38f-ab1537f889cf MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:42 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:42.502105991 +0300 MSK m=+36.002297425 MessageId:e3792aa1-8da4-4461-a4d5-768e490fe53b MessageText:EFGHIJKLMNOPQRSTUVWX}
2025/04/28 07:55:42 <<<Reply message: {MessageId:e3792aa1-8da4-4461-a4d5-768e490fe53b MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:43 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:43.502381928 +0300 MSK m=+37.002573363 MessageId:1c20cd0f-7baa-418b-b5eb-db3484d1682d MessageText:YZ0123456789bcdefghi}
2025/04/28 07:55:43 <<<Reply message: {MessageId:1c20cd0f-7baa-418b-b5eb-db3484d1682d MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:44 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:44.50247194 +0300 MSK m=+38.002663403 MessageId:7f3089ce-6ce6-43ac-b0f8-dbe465f47a09 MessageText:jklmnopqrstuvwxyzABC}
2025/04/28 07:55:44 <<<Reply message: {MessageId:7f3089ce-6ce6-43ac-b0f8-dbe465f47a09 MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:45 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:45.502093636 +0300 MSK m=+39.002285102 MessageId:de87b373-7061-4d67-94f3-8fd5a41d06b0 MessageText:DEFGHIJKLMNOPQRSTUVW}
2025/04/28 07:55:45 <<<Reply message: {MessageId:de87b373-7061-4d67-94f3-8fd5a41d06b0 MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:46 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:46.502082396 +0300 MSK m=+40.002273853 MessageId:0495f55d-ec3c-4ea4-9acc-9f530eea8d01 MessageText:XYZ0123456789bcdefgh}
2025/04/28 07:55:46 <<<Reply message: {MessageId:0495f55d-ec3c-4ea4-9acc-9f530eea8d01 MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:47 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:47.501955188 +0300 MSK m=+41.002146680 MessageId:278b0bfe-b96c-4686-8a5c-cb677f0dffe7 MessageText:ijklmnopqrstuvwxyzAB}
2025/04/28 07:55:47 <<<Reply message: {MessageId:278b0bfe-b96c-4686-8a5c-cb677f0dffe7 MessageStatus:200 MessageReply:OK from NN}
2025/04/28 07:55:48 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:48.501942683 +0300 MSK m=+42.002134120 MessageId:5bc77854-5107-4742-905e-830e2d2c0a5f MessageText:CDEFGHIJKLMNOPQRSTUV}
2025/04/28 07:55:48 <<<Reply message: {MessageId:5bc77854-5107-4742-905e-830e2d2c0a5f MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:49 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:49.502427218 +0300 MSK m=+43.002618652 MessageId:1881d4e2-8e82-4388-b624-53198ed0c5c4 MessageText:WXYZ0123456789bcdefg}
2025/04/28 07:55:49 <<<Reply message: {MessageId:1881d4e2-8e82-4388-b624-53198ed0c5c4 MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:50 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:50.502526369 +0300 MSK m=+44.002717861 MessageId:f30f93e1-a06e-4ff7-bf10-0680d73cedde MessageText:hijklmnopqrstuvwxyzA}
2025/04/28 07:55:50 <<<Reply message: {MessageId:f30f93e1-a06e-4ff7-bf10-0680d73cedde MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:51 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:51.502245904 +0300 MSK m=+45.002437396 MessageId:90ad410e-2dff-4bd7-9221-d339f8176e56 MessageText:BCDEFGHIJKLMNOPQRSTU}
2025/04/28 07:55:51 <<<Reply message: {MessageId:90ad410e-2dff-4bd7-9221-d339f8176e56 MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:52 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:52.502418873 +0300 MSK m=+46.002610308 MessageId:19d7b26e-d544-4c43-83a6-45fee984da02 MessageText:VWXYZ0123456789bcdef}
2025/04/28 07:55:52 <<<Reply message: {MessageId:19d7b26e-d544-4c43-83a6-45fee984da02 MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:53 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:53.502504033 +0300 MSK m=+47.002695467 MessageId:d26b2382-3754-4829-afc7-811dc935cec6 MessageText:ghijklmnopqrstuvwxyz}
2025/04/28 07:55:53 <<<Reply message: {MessageId:d26b2382-3754-4829-afc7-811dc935cec6 MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:54 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:54.502493987 +0300 MSK m=+48.002685421 MessageId:e8971a31-ab4d-4940-a3e2-e1703e0e6fc1 MessageText:ABCDEFGHIJKLMNOPQRST}
2025/04/28 07:55:54 <<<Reply message: {MessageId:e8971a31-ab4d-4940-a3e2-e1703e0e6fc1 MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:55 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:55.502447344 +0300 MSK m=+49.002638837 MessageId:1c8c6334-5002-4930-89bd-e9f493a545f2 MessageText:UVWXYZ0123456789bcde}
2025/04/28 07:55:55 <<<Reply message: {MessageId:1c8c6334-5002-4930-89bd-e9f493a545f2 MessageStatus:200 MessageReply:OK from YAR}
2025/04/28 07:55:56 >>>[YAR] Published message: {Timestamp:2025-04-28 07:55:56.502624397 +0300 MSK m=+50.002815941 MessageId:0f637835-acb6-48f9-8033-7a8605bf10c8 MessageText:fghijklmnopqrstuvwxy}
2025/04/28 07:55:56 <<<Reply message: {MessageId:0f637835-acb6-48f9-8033-7a8605bf10c8 MessageStatus:200 MessageReply:OK from YAR}
```

#### 3. Work Queue (group consumers)


Запустить consumer в количестве трех штук
go run consumer-queue/consumer.go -group group1 -url nats://localhost:4222

Пример журнала
```text
2025/04/28 08:18:14 Consumer connected to NATS  YAR
2025/04/28 08:18:14 Queue group group1
Received message: &{Subject:events Reply:_INBOX.1AdlXZ6Mc7IBKc57mOY1zo.z7YPWtK3 Header:map[] Data:[123 34 116 105 109 101 115 116 97 109 112 34 58 34 50 48 50 53 45 48 52 45 50 56 84 48 56 58 49 56 58 49 53 46 54 51 48 52 57 51 54 49 56 43 48 51 58 48 48 34 44 34 109 101 115 115 97 103 101 95 105 100 34 58 34 56 99 53 48 51 48 102 54 45 49 98 101 100 45 52 54 53 57 45 97 98 56 101 45 102 100 57 51 99 51 98 53 101 57 98 99 34 44 34 109 101 115 115 97 103 101 95 116 101 120 116 34 58 34 122 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 34 125] Sub:0xc000272000 next:<nil> wsz:185 barrier:<nil> ackd:0}
Received message: &{Subject:events Reply:_INBOX.1AdlXZ6Mc7IBKc57mOY1zo.cnvkry4z Header:map[] Data:[123 34 116 105 109 101 115 116 97 109 112 34 58 34 50 48 50 53 45 48 52 45 50 56 84 48 56 58 49 56 58 49 54 46 54 51 48 52 49 49 55 57 57 43 48 51 58 48 48 34 44 34 109 101 115 115 97 103 101 95 105 100 34 58 34 57 54 51 100 100 100 49 52 45 50 48 51 99 45 52 52 99 53 45 97 51 54 57 45 55 98 99 56 98 101 99 52 97 55 98 100 34 44 34 109 101 115 115 97 103 101 95 116 101 120 116 34 58 34 84 85 86 87 88 89 90 48 49 50 51 52 53 54 55 56 57 98 99 100 34 125] Sub:0xc000272000 next:<nil> wsz:185 barrier:<nil> ackd:0}
^C2025/04/28 08:18:18 Shutting down consumer...
```

Запустить producer 

```shell
go run producer/producer.go -url nats://localhost:4222
```
Журнал работы
```text
2025/04/28 08:21:17 Producer connected to NATS  YAR
2025/04/28 08:21:18 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:18.366724632 +0300 MSK m=+1.002906045 MessageId:3aec4496-0a04-4408-a208-6ead154a947c MessageText:bcdefghijklmnopqrstu}
2025/04/28 08:21:18 <<<Reply message: {MessageId:3aec4496-0a04-4408-a208-6ead154a947c MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 51}
2025/04/28 08:21:19 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:19.366495767 +0300 MSK m=+2.002677186 MessageId:8382546e-5b80-4476-acff-f9329504f1b2 MessageText:vwxyzABCDEFGHIJKLMNO}
2025/04/28 08:21:19 <<<Reply message: {MessageId:8382546e-5b80-4476-acff-f9329504f1b2 MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 51}
2025/04/28 08:21:20 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:20.366531086 +0300 MSK m=+3.002712442 MessageId:82f043bd-f5df-4c27-95dc-96e5db3e70b7 MessageText:PQRSTUVWXYZ012345678}
2025/04/28 08:21:20 <<<Reply message: {MessageId:82f043bd-f5df-4c27-95dc-96e5db3e70b7 MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 52}
2025/04/28 08:21:21 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:21.366554173 +0300 MSK m=+4.002735567 MessageId:bdab7d6f-a726-4c68-ab94-fd601c65a6f2 MessageText:9bcdefghijklmnopqrst}
2025/04/28 08:21:21 <<<Reply message: {MessageId:bdab7d6f-a726-4c68-ab94-fd601c65a6f2 MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 52}
2025/04/28 08:21:22 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:22.366350734 +0300 MSK m=+5.002532006 MessageId:811f1e8b-f30a-4e76-933f-68266c3dd3ce MessageText:uvwxyzABCDEFGHIJKLMN}
2025/04/28 08:21:22 <<<Reply message: {MessageId:811f1e8b-f30a-4e76-933f-68266c3dd3ce MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 52}
2025/04/28 08:21:23 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:23.366513744 +0300 MSK m=+6.002695099 MessageId:5c55d23f-f38a-41c9-84cc-90013f956543 MessageText:OPQRSTUVWXYZ01234567}
2025/04/28 08:21:23 <<<Reply message: {MessageId:5c55d23f-f38a-41c9-84cc-90013f956543 MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 53}
2025/04/28 08:21:24 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:24.366100623 +0300 MSK m=+7.002281991 MessageId:851a3905-1424-45d2-a358-aa23d6fca907 MessageText:89bcdefghijklmnopqrs}
2025/04/28 08:21:24 <<<Reply message: {MessageId:851a3905-1424-45d2-a358-aa23d6fca907 MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 51}
2025/04/28 08:21:25 >>>[YAR] Published message: {Timestamp:2025-04-28 08:21:25.366744675 +0300 MSK m=+8.002926085 MessageId:ed39410f-dcdd-48b8-b84c-771474741867 MessageText:tuvwxyzABCDEFGHIJKLM}
2025/04/28 08:21:25 <<<Reply message: {MessageId:ed39410f-dcdd-48b8-b84c-771474741867 MessageStatus:200 MessageReply:OK from YAR group group1 consumer id 53}
```
Можно обратить внимание, что в ответе в поле MessageReply что ответы поступают от разных consumer id.
