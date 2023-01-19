# REST API Для Создания TODO Списков на Go

## Цель этого проекта:
- Разработка Веб-Приложения на Go, следуя дизайну REST API
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости
- Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>
- Работа с переменными окружения
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>
- Регистрация и аутентификация. Работа с JWT. Middleware
- Написание SQL запросов
- Graceful Shutdown

## Как затестить

### Для запуска БД:

```
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p=5432:5432 -d --rm postgres
```

### Для инициализации БД:
установить <a href="https://github.com/golang-migrate/migrate/tree/master/cmd/migrate">migrate</a>
```
sudo migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

### Для запуска приложения:

```
go run cmd/main.go
```

### Команды для постмана
```
POST localhost:8000/auth/sign-up
```
в raw в формате JSON
```
{
    "name":"qwe",
    "username": "asd",
    "password": "qwerty"
}
```
```
POST localhost:8000/auth/sign-in
```
в raw в формате JSON
```
{
    "username": "asd",
    "password": "qwerty"
}
```
Тут в ответе будет токен, его нужно будет использовать везде дальше. Во всех следующих командах
в Authorithation выбрать Bearer Token, скопировать в поле token полученный token
```
POST localhost:8000/api/lists
```
в raw в формате JSON
```
{
    "title": "План дня",
    "description": "Дела"
}
```
```
GET localhost:8000/api/lists
```
```
GET localhost:8000/api/lists/1
```
```
PUT localhost:8000/api/lists/1
```
в raw в формате JSON
```
{
    "description":"купить сегодня"
}
```
```
POST localhost:8000/api/lists/1/items
```
в raw в формате JSON
```
{
    "title": "qwerty",
    "description": "asdfgh"
}
```
```
GET localhost:8000/api/lists/1/items
```
```
PUT localhost:8000/api/items/1
```
в raw в формате JSON
```
{
    "done":true
}
```
```
GET localhost:8000/api/items/1
```
```
DELETE localhost:8000/api/items/1
```
```
DELETE localhost:8000/api/lists/1
```
