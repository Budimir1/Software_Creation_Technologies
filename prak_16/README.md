# **Практическое занятие №16: Интеграционное тестирование API. Использование Docker для тестовой БД**

**ФИО:** Булискерия Будимир Бакуриевич  
**Группа:** ПИМО-01-25  
**Дисциплина:** Технологии создания программного обеспечения  
**Институт:** ПИШ  
**Кафедра:** Кафедра передовых технологий  
**Преподаватель:** Адышкин Сергей Сергеевич  
**Семестр:** 1 семестр, 2025-2026 гг.

---

## Описание

Проект посвящён изучению интеграционного тестирования REST API с использованием реальной базы данных в Docker.  
В ходе работы реализованы два подхода к организации тестовой среды:

1. **Вариант A:** Локальная среда через `docker-compose`
2. **Вариант B:** Программный подъём контейнеров через `testcontainers-go`

---

## Структура проекта

pz16-integration/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── models/
│   │   └── note.go
│   ├── repo/
│   │   ├── postgres.go
│   │   └── user_repo.go
│   ├── service/
│   │   └── service.go
│   ├── httpapi/
│   │   └── handlers.go
│   └── db/
│       └── migrate.go
├── integration/
│   ├── notes_integration_test.go      # Вариант A
│   └── notes_tc_integration_test.go   # Вариант B
├── testdata/
│   └── fixtures/
├── docker-compose.yml                 # Вариант A
├── Makefile
├── go.mod
└── README.md

---

## Модель данных

```go
type Note struct {
    ID        int64  `db:"id" json:"id"`
    Title     string `db:"title" json:"title"`
    Content   string `db:"content" json:"content"`
    CreatedAt string `db:"created_at" json:"created_at"`
    UpdatedAt string `db:"updated_at" json:"updated_at"`
}
```
---

## Реализованные эндпоинты

- POST /notes – создание заметки
- GET /notes/:id – получение заметки по ID
- PUT /notes/:id – обновление заметки
- DELETE /notes/:id – удаление заметки
- GET /notes – список заметок (опционально)

---

## Инструкции по запуску

### Предварительные требования
- Docker и Docker Compose
- Go 1.21+

### Вариант A: Docker Compose

1. Клонирование и инициализация
mkdir prak_16 && cd prak_16
go mod init Budimir/prak_16

2. Запуск тестовой базы данных
docker-compose up -d

3. Запуск интеграционных тестов
DB_DSN="postgres://test:test@localhost:54321/notes_test?sslmode=disable" go test -v ./integration/

4. Остановка базы данных
docker-compose down -v

### Вариант B: Testcontainers-go

1. Установка зависимостей
go get github.com/testcontainers/testcontainers-go@latest
go get github.com/testcontainers/testcontainers-go/modules/postgres@latest

2. Запуск тестов
go test -v ./integration/

---
## Проверка
![img.png](image%2Fimg.png)
![img_1.png](image%2Fimg_1.png)
## Примеры успешного вывода тестов:
![img_2.png](image%2Fimg_2.png)
![img_3.png](image%2Fimg_3.png)
![img_4.png](image%2Fimg_4.png)