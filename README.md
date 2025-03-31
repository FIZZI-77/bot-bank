
### 1️⃣ Запуск PostgreSQL в контейнере
```sh
docker run --name=telegram-db -e POSTGRES_DB=tg-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
```

### 2️⃣ Выполнение миграций (создание БД)
```sh
goose -dir src/scheme  postgres "postgresql://postgres:qwerty@127.0.0.1:5436/tg-db?sslmode=disable" up
```

### 3️⃣ Запуск самого приложения
```sh
go run cmd/main.go
```