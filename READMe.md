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

## Формат запросов

### HTTP 

#### Сохраняет URL в хранилище и возвращает сокращенный

* Метод: `POST`
* Эндпоинт: `http://localhost:8080/shorten`
* Формат тела запроса:

```json
{
    "url": "https://www.youtube.com"
}
```

* Формат ответа:

```json
{
  "alias": "SL9bwykhFs"
}
```

#### Принимает сокращенный URL, возвращает оригинальный 

* Метод: `GET`
* Эндпоинт: `http://localhost:8080/get-url`
* Формат тела запроса:

```json
{
    "alias": "SL9bwykhFs"
}
```

* Формат ответа:

```json
{
  "alias": "https://www.youtube.com"
}
```

### GRPC

#### Сохраняет URL в хранилище и возвращает сокращенный

* Обработчик: `ShortenUrl`
* Формат запроса и ответа:

```text
app.UrlShortener@127.0.0.1:8080> call ShortenUrl
url (TYPE_STRING) => http://localhost
{                      
  "alias": "CqTJate_ol"
}   
```

#### Принимает сокращенный URL, возвращает оригинальный

* Обработчик: `GetUrl`
* Формат запроса и ответа:

```text
app.UrlShortener@127.0.0.1:8080> call GetUrl
alias (TYPE_STRING) => CqTJate_ol
{
  "url": "http://localhost"
}
 
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
