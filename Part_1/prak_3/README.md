# Практическое занятие №3: HTTP-сервер на Go

## Описание
Реализация простого HTTP-сервера для управления задачами с использованием стандартной библиотеки net/http.

## Структура проекта
```azure
pz3-http/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers.go
│   │   ├── middleware.go
│   │   └── responses.go
│   └── storage/
│       └── memory.go
├── go.mod
├── requests.md
└── README.md
```

## Функциональность
- ✅ GET /health - проверка состояния сервера
- ✅ GET /tasks - получение списка задач с фильтрацией
- ✅ POST /tasks - создание новой задачи
- ✅ GET /tasks/{id} - получение задачи по ID
- ✅ PATCH /tasks/{id} - обновление задачи
- ✅ DELETE /tasks/{id} - удаление задачи
- ✅ Middleware для логирования
- ✅ Middleware для CORS
- ✅ Валидация данных (длина title: 3-140 символов)
- ✅ Graceful shutdown
- ✅ Обработка ошибок с соответствующими HTTP-статусами

## Запуск проекта

### Предварительные требования
- Go 1.21 или выше
- Утилиты для тестирования: curl, Postman или аналоги

### Установка и запуск

1. **Инициализация проекта:**
```bash
mkdir -p pz3-http/cmd/server pz3-http/internal/api pz3-http/internal/storage
cd pz3-http
go mod init example.com/pz3-http