# TechIP Tasks Service

Микросервис для управления задачами, созданный в рамках учебной практики по CI/CD (Технологии создания программного обеспечения).

## Сервис

- **Язык**: Go 1.23+
- **HTTP API**: порт 8080
- **Эндпоинты**:
    - `GET /health` – проверка работоспособности
    - `GET /tasks` – список задач
    - `POST /tasks` – добавить задачу (JSON: `{"title":"..."}`)
- **Хранение**: в памяти (in-memory)

## Локальный запуск (без Docker)

```bash
cd services/tasks
go mod tidy
go run .
```
Сервис будет доступен на http://localhost:8080.
Пример проверки:

```bash
curl http://localhost:8080/health
```
Запуск тестов
```bash
cd services/tasks
go test ./...
```
Docker (опционально, локально или в CI)
Сборка образа:

```bash
cd services/tasks
docker build -t techip-tasks:0.1 .
```
Запуск:

```bash
docker run -p 8080:8080 techip-tasks:0.1
CI/CD
```
В проекте настроен pipeline на GitHub Actions:

запуск тестов и сборки Go

сборка Docker-образа



Конфигурация: .github/workflows/ci.yml

Структура проекта
```text
.
└── services/
└── tasks/
├── main.go
├── go.mod / go.sum
├── Dockerfile
├── .dockerignore
└── handlers/
├── handlers.go
└── handlers_test.go
```