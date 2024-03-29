# 💼 Тестовое задание

**🐾** Задача: разработать бэкенд приложения "фильмотека", который представляет REST API для управления базой данных фильмов.

## 😎 **Swagger**

http://localhost:9091/swagger/index.html

## 📜 Справка

📌 Порты у БД открыты для дебага

📌 Покрытие тестами можно посмотреть через `go test -cover ./... `. Однако при запуске из корневой директории отображаются неверные проценты

📌 В редисе написаны методы для обновления рефреш токенов и акссес токенов, но в ручках реализации еще нет

📌 Все эндпоинты закрыты от guest\`ов. Для user\`ов только получение данных и авторизация

## 🩻 Структура проекта
- cmd/api - _**main.go**_
- config/ - _yaml и структура_
- data/
    - postgres
    - redis
- internal
    - api - _реализация хендлеров в трехслойной архитектуре_
      - delivery - _хендлеры_
      - repository - _взаимодействие с **PostgreSQL** и **Redis** (хранение refresh token)_
      - usecase - _бизнес-логика_
      - ========================================
      - handler.go -          _интерфейс хенделера_
      - repository.go -     _интерфейс репозитория_
      - tokenRepository.go -_интерфейс репозитория токенов_
      - usecase.go - _интерфейс юзкейса)_  
    - common/constant - _константы_ 
    - middleware - _middleware для авторизации_
    - server
        - delivery/mapRoutes - _рутинг_
        - mapHandlers.go - _инициализация инстансов_
        - runServer.go - _запуск сервера_
    - utils/encryption - _хеширование пароля_
- migrations - _sql миграции_
- pkg/logger - _логгер_
## 🧶 Config sample

```
Server:
  Port: 0000
  Version: 1.0.0
  AccessSecret: secret
  RefreshSecret: rsecret
  AccessLifetime: 7200
  RefreshLifetime: 604800

Logger:
  InFile: true

Postgres:
  Host: cinema_db
  Port: 0000
  User: postgres_user_from_env
  Password: postgres_password_from_env
  Database: cinema
  SSLMode: disable

Redis:
  Host: api_redis_db
  Port: 0000
  Database: 0
```

## 🐈 .env file sample

```
POSTGRES_USER=postgres_user
POSTGRES_PASSWORD=postgres_password
POSTGRES_DB=cinema
DEBUG_LEVEL=0

IN_DOCKER=TRUE
```