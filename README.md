# Order Service
This project is a demo Golang service for displaying order data. It uses PostgreSQL, data streaming and caching to optimize data access.
## Содержание
- [Технологии](#технологии)
- [Начало работы](#начало-работы)
- [Тестирование](#тестирование)
- [Deploy и CI/CD](#deploy-и-ci/cd)
- [Contributing](#contributing)
- [To do](#to-do)
- [Команда проекта](#команда-проекта)

## Технологии
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [NATS-Streaming](https://github.com/nats-io/nats-streaming-server)
- [pgx](https://github.com/jackc/pgx)
- [gin](https://github.com/gin-gonic/gin)
- [Swagger](https://swagger.io/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Vegeta](https://github.com/tsenart/vegeta)
- [go-wrk](https://github.com/tsliwowicz/go-wrk)

## Использование
1. Clone this repository

Установите npm-пакет с помощью команды:
```sh
$ npm i your-awesome-plugin-name
```

И добавьте в свой проект:
```typescript
import { hi } from "your-awesome-plugin-name";

hi(); // Выведет в консоль "Привет!"
```

## Разработка

### Требования
Для установки и запуска проекта, необходим [NodeJS](https://nodejs.org/) v8+.

### Установка зависимостей
Для установки зависимостей, выполните команду:
```sh
$ npm i
```

### Запуск Development сервера
Чтобы запустить сервер для разработки, выполните команду:
```sh
npm start
```

### Создание билда
Чтобы выполнить production сборку, выполните команду: 
```sh
npm run build
```

## Тестирование
Какие инструменты тестирования использованы в проекте и как их запускать. Например:

Наш проект покрыт юнит-тестами Jest. Для их запуска выполните команду:
```sh
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

### Зачем вы разработали этот проект?
Чтобы был.

## To do
- [ ] Добавить больше тестов
- [ ] Добавить больше функционала
- [ ] Сделать полноценный UI интерфейс с большим функционалом
