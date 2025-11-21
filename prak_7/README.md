# Практическое занятие №7 - Redis кэширование

Простой HTTP-сервер на Go с реализацией кэширования данных с использованием Redis.

## Описание

Проект демонстрирует базовые операции работы с Redis из Go-приложения:
- Сохранение и получение данных (SET/GET)
- Установка времени жизни ключей (TTL)
- Реализация простого кэша для ускорения работы API

## Технологии

- **Go** - язык программирования
- **Redis** - in-memory key-value хранилище
- **Docker** - для запуска Redis

## Требования

- Установленный Go (версия 1.16 или выше)
- Установленный Docker
- Доступ в интернет для загрузки зависимостей

## Установка и запуск

### 1. Клонирование и инициализация проекта

```bash
mkdir prak_7
cd prak_7
go mod init Budimir/prak_7
go get github.com/redis/go-redis/v9
```
## Запуск Redis
```bash
docker run --name redis -p 6379:6379 redis
```
## Запуск приложения
```bash
go run ./cmd/server
```
## API Endpoints
### Установка значения
```
GET /set?key=<ключ>&value=<значение>
```
### Получение значения
```
GET /get?key=<ключ>
```
### Проверка TTL
```
GET /ttl?key=<ключ>
```
## Примеры использования
```bash
# Сохранить значение
curl "http://localhost:8080/set?key=username&value=john_doe"

# Получить значение
curl "http://localhost:8080/get?key=username"

# Проверить TTL
curl "http://localhost:8080/ttl?key=username"
```
## Структура проекта
```
pz7-redis/
├── cmd/
│   └── server/
│       └── main.go          # Основной HTTP-сервер
├── internal/
│   └── cache/
│       └── cache.go         # Реализация кэша
├── go.mod
└── README.md
```