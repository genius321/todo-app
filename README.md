# REST API Для Создания TODO Списков на Go

## Цель этого проекта:
- Разработка Веб-Приложения на Go, следуя дизайну REST API.
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
- Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций. 
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>
- Работа с переменными окружения.
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Регистрация и аутентификация. Работа с JWT. Middleware.
- Написание SQL запросов.
- Graceful Shutdown

## Usage

### Для запуска БД:

```
docker pull postgres
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p=5432:5432 -d --rm postgres
```

### Для инициализации БД:
```
sudo migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```

### Для запуска приложения:

```
go run cmd/main.go
```