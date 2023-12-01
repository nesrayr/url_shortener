# url_shortener

## Описание 

Сервис, предоставляющий API по созданию сокращенных ссылок. 

## Структура 

```text
├───cmd
│   └───app
│           main.go
│
├───config
│       config.go
│
├───internal
│   ├───migrations
│   │       20231129173251_create_tables.go
│   │       migrations.go
│   │
│   ├───ports
│   │   ├───grpc
│   │   │       handler.go
│   │   │       server.go
│   │   │
│   │   └───http
│   │           handler.go
│   │           router.go
│   │
│   ├───repo
│   │   │   repo.go
│   │   │
│   │   ├───cache
│   │   │       repo.go
│   │   │
│   │   ├───mocks
│   │   │       mock_repo.go
│   │   │
│   │   └───postgres
│   │           queries.go
│   │           repo.go
│   │
│   ├───service
│   │       coverage.out
│   │       service.go
│   │       service_test.go
│   │       utils.go
│   │       utils_test.go
│   │
│   └───storage
│       ├───cache
│       │       storage.go
│       │
│       └───postgres
│               config.go
│               storage.go
│
├───pkg
│   └───logging
│           logging.go
│
└───protos
    ├───gen
    │       app.pb.go
    │       app_grpc.pb.go
    │
    └───proto
            app.proto
```

## Запуск 

Сначала необходимо добавить конфигурационный файл `.env` в корне проекта. Пример конфигурации:

```text
PORT=8080
HOST=localhost

#db
DB_USER=user
DB_PASSWORD=password
DB_NAME=url
DB_HOST=postgres
DB_POOL_SIZE=100
DB_PORT=5432
```

Сборка и запуск контейнеров:

```bash
make up-http-postgres - # Запуск http-сервера, postgresql - в качестве хранилища
```
```bash
make up-http-cache - # Запуск http-сервера, in-memory cache - в качестве хранилища
```
```bash
make up-grpc-postgres - # Запуск grpc-сервера, postgresql - в качестве хранилища
```
```bash
make up-grpc-cache - # Запуск grpc-сервера, in-memory cache - в качестве хранилища
```

## Тестирование

Для запуска тестов:

```bash
make test
```

Для просмотра test coverage: 

```bash
make cover
```
