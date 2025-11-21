# CRUD-сервис "Список задач" (ToDo)

Проект представляет собой REST API для управления списком задач, реализованный на Go с использованием роутера chi.

## Цели проекта

1. Освоить базовую маршрутизацию HTTP-запросов в Go с помощью роутера chi
2. Научиться строить REST-маршруты и обрабатывать методы GET/POST/PUT/DELETE
3. Реализовать CRUD-сервис для управления задачами с хранением в памяти
4. Добавить middleware для логирования и CORS
5. Научиться тестировать API через curl/Postman/HTTPie

## Технологии

- Go 1.21
- Chi router v5
- Стандартная библиотека Go

## Структура проекта
```azure
pz4-todo/
├── internal/
│   └── task/
│       ├── model.go
│       ├── repo.go
│       └── handler.go
├──  pkg/
│    └── middleware/
│        └──logger.go
│        └── cors.go
├── go.mod
├── go.sum
└── main.go

```

## Установка и запуск

1. Клонируйте репозиторий:
```azure
git clone <repository-url>
cd pz4-todo 
``` 

2. Установите зависимости:
```azure
go mod download
```
3. Запустите сервер:
 ```azure
go run main.go
```

Сервер будет доступен по адресу: http://localhost:8080

# API Endpoints
## Проверка здоровья сервера

```azure
GET /health
```
Response: OK

## Работа с задачами
Получить список всех задач
```azure 
GET /api/tasks
```
Response: 200 OK с массивом задач

Создать новую задачу
```azure
POST /api/tasks
Content-Type: application/json

{
  "title": "Название задачи"
}
```
Response: 201 Created с созданной задачей

Получить задачу по ID
```azure
GET /api/tasks/{id}
```
Response: 200 OK с задачей или 404 Not Found

Обновить задачу
```azure
PUT /api/tasks/{id}
Content-Type: application/json

{
  "title": "Обновленное название",
  "done": true
}
```
Response: 200 OK с обновленной задачей или 404 Not Found

Удалить задачу
```azure
DELETE /api/tasks/{id}
```
Response: 204 No Content или 404 Not Found

# Примеры использования
## Создание задачи
```azure
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Изучить Go"}'
```
Получение списка задач
```azure
curl http://localhost:8080/api/tasks
```

Получение задачи по ID
```azure
curl http://localhost:8080/api/tasks/1
```
Обновление задачи   
```azure
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Изучить Go глубже","done":true}'
```
Удаление задачи
```azure
curl -X DELETE http://localhost:8080/api/tasks/1
```
Модель данных
```azure
type Task struct {
    ID        int64     `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```
