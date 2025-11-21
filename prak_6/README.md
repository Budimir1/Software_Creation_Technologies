# Практическое занятие №6: Использование GORM

Проект демонстрирует использование ORM GORM в Go для работы с моделями, миграциями и связями между таблицами в PostgreSQL.

## Описание

GORM позволяет работать с базой данных через Go-структуры, автоматически создавать таблицы и управлять связями между ними (1:N, M:N). Этот проект включает минимальный REST API для работы с пользователями и заметками.

## Требования

- Go 1.21 и выше
- PostgreSQL 18

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone <ваш-репозиторий>
cd pak_6
```
## Установите зависимости:
```
go mod tidy
```
## Создайте базу данных:
```
CREATE DATABASE postgres;
```
## API Endpoints
### Health Check
```
curl http://localhost:8080/health
```
### Создание пользователя
```
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```
### Создание заметки
```
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title":"Первая заметка","content":"Текст...","userId":1,"tags":["go","gorm"]}'
  ```
### Получение заметки
```
curl http://localhost:8080/notes/1
```
## Структура проекта
```
pz6-gorm/
├── cmd/server/main.go          # Точка входа
├── internal/
│   ├── db/postgres.go          # Подключение к БД
│   ├── models/models.go        # Модели GORM
│   └── httpapi/
│       ├── router.go           # Роутер Chi
│       └── handlers.go         # Обработчики API
└── go.mod
```