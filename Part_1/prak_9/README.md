# Реализация регистрации и входа пользователей. Хеширование gfhjktq c bcrypt

Минимальный сервис для реализации регистрации и входа пользователей с безопасным хэшированием паролей.

## Описание

Проект реализует два основных эндпоинта:
- `POST /auth/register` - регистрация нового пользователя
- `POST /auth/login` - аутентификация существующего пользователя

Пароли хэшируются с использованием алгоритма bcrypt с автоматической генерацией соли.

## Технологии

- **Go** (Golang)
- **PostgreSQL** - база данных
- **GORM** - ORM для работы с БД
- **Chi** - маршрутизатор HTTP
- **bcrypt** - хэширование паролей

## Структура проекта
```
pz9-auth/
├── cmd/
│ └── api/
│ └── main.go # Точка входа приложения
├── internal/
│ ├── core/
│ │ └── user.go # Модель пользователя
│ ├── http/
│ │ └── handlers/
│ │ └── auth.go # HTTP-обработчики
│ ├── repo/
│ │ ├── postgres.go # Подключение к БД
│ │ └── user_repo.go # Репозиторий пользователей
│ └── platform/
│ └── config/
│ └── config.go # Конфигурация
└── go.mod
```

# Установка и запуск

## 1. Клонирование и инициализация

```bash
mkdir prak_9 && cd prak_9
go mod init Budimir/prak_9
```
###  Установка зависимостей
```bash
go get github.com/go-chi/chi/v5
go get gorm.io/gorm gorm.io/driver/postgres
go get golang.org/x/crypto/bcrypt
```
## Настройка окружения
### Установите переменные окружения:
```powershell
export DB_DSN="postgres://user:pass@localhost:5432/pz9?sslmode=disable"
export BCRYPT_COST=12
export APP_ADDR=":8080"
```
## Использование API
### Регистрация пользователя
```bash
curl -i -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Secret123!"}'
```
### Вход пользователя
```bash
curl -i -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Secret123!"}'
```