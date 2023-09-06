# Order Service
This project is a demo Golang service for displaying order data. It uses PostgreSQL, data streaming and caching to optimize data access.
## Content
- [Technologies](#technologies)
- [Usage](#usage)
- [Request examples](#request-examples)
- [Stress test](#stress-test)
- [Why did you develop this project?](#why-did-you-develop-this-project?)
- [To do](#to-do)

## Technologies
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [NATS-Streaming](https://github.com/nats-io/nats-streaming-server)
- [pgx](https://github.com/jackc/pgx)
- [gin](https://github.com/gin-gonic/gin)
- [Swagger](https://swagger.io/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Vegeta](https://github.com/tsenart/vegeta)
- [go-wrk](https://github.com/tsliwowicz/go-wrk)

## Usage
1. Clone this repository
```sh
git clone https://github.com/EldarSahipov/OrderService.git
```
2. Go to folder
```
go to folder OrderService
```
3. Build the project with docker compose and run
```
docker compose build && docker compose up -d 
```
4. Since the database will be empty, it is necessary to migrate the tables from the folder: ```./schema/```
```
migrate -path ./schema -database "postgres://postgres:1234@localhost:5432/order_management_db?sslmode=disable" up
```

## Request examples
### GET /api/orders/b563feb7b2b84b6test
```
{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```

### GET /api/orders/all
```
[
  {
    "order_uid": "b563feb7b2b84b6test",
    "track_number": "WBILMTESTTRACK",
    "entry": "WBIL",
    "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
    },
    "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
    },
    "items": [
      {
        "chrt_id": 9934930,
        "track_number": "WBILMTESTTRACK",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "Mascaras",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
      }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
  }
]
```
### Bad Request
```
{
  "message": "no rows in result set"
}
```

## Stress test
### 1. Vegeta.
#### Testing method: ```/api/orders/{uid}```
```
PS C:\dev\Wildberries\OrderService> echo "GET http://localhost:8080/api/orders/b563feb7b2b84b6test" | vegeta attack -duration=5s -rate=200/s --output results.bin | vegeta report results.bin
Requests      [total, rate, throughput]  1000, 200.30, 200.28
Duration      [total, attack, wait]      4.9930248s, 4.9924414s, 583.4µs
Latencies     [mean, 50, 95, 99, max]    739.182µs, 670.977µs, 1.12621ms, 1.38345ms, 15.103ms
Bytes In      [total, mean]              835000, 835.00
Bytes Out     [total, mean]              0, 0.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:1000
Error Set:
```
### 2. go-wrk
#### Testing method: ```/api/orders/{uid}```
```
Running 5s test @ http://localhost:8080/api/orders/b563feb7b2b84b6test
  80 goroutine(s) running concurrently
98561 requests in 4.849667817s, 88.54MB read
Requests/sec:           20323.25
Transfer/sec:           18.26MB
Avg Req Time:           3.936378ms
Fastest Request:        0s
Slowest Request:        354.4799ms
Number of Errors:       0
```

## Why did you develop this project?
To be.

## To do
- [ ] Add more tests
- [ ] Add more functionality
- [ ] Make a full-fledged UI interface with great functionality
