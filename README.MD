# Тестовое задание

## Структура проекта
- cmd/api - **main.go**
- config/ - yaml и структура
- data/
    - postgres
    - redis
- internal
    - api - реализация хендлеров в трехслойной архитектуре
      - delivery - хендлеры
      - repository - взаимодействие с **PostgreSQL** и **Redis** (хранение refresh token)
      - usecase - бизнес-логика
      - ========================================
      - handler.go -          интерфейс хенделера
      - repository.go -     интерфейс репозитория
      - tokenRepository.go -интерфейс репозитория токенов
      - usecase.go - интерфейс юзкейса)  
    - common/constant - константы 
    - middleware - middleware для авторизации
    - server
        - delivery/mapRoutes - рутинг
        - mapHandlers.go - инициализация инстансов
        - runServer.go - запуск сервера
    - utils/encryption - хеширование пароля
- migrations - sql миграции
- pkg/logger - логгер
## Config sample

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

## .env file sample

```
POSTGRES_USER=postgres_user
POSTGRES_PASSWORD=postgres_password
POSTGRES_DB=cinema
DEBUG_LEVEL=0

IN_DOCKER=TRUE
```