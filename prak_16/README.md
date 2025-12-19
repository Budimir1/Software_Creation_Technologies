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
mkdir pz16-integration && cd pz16-integration
go mod init example.com/pz16-integration

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

## Docker Compose конфигурация

version: "3.9"
services:
  pg:
    image: postgres:16
    environment:
      POSTGRES_USER: test
      POSTGRES_PASSWORD: test
      POSTGRES_DB: notes_test
    ports:
      - "54321:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U test -d notes_test"]
      interval: 2s
      timeout: 2s
      retries: 20

---

## Пример интеграционного теста

func TestCreateAndGetNote(t *testing.T) {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        t.Skip("DB_DSN not set")
    }
    
    srv := newServer(t, dsn)
    defer srv.Close()
    
    // Создание заметки
    resp, err := http.Post(srv.URL+"/notes", "application/json",
        strings.NewReader(`{"title":"Hello","content":"World"}`))
    if err != nil {
        t.Fatal(err)
    }
    
    if resp.StatusCode != http.StatusCreated {
        t.Fatalf("status %d != 201", resp.StatusCode)
    }
    
    // Проверка получения заметки
    var created map[string]any
    body, _ := io.ReadAll(resp.Body)
    _ = json.Unmarshal(body, &created)
    id := int64(created["id"].(float64))
    
    resp2, err := http.Get(fmt.Sprintf("%s/notes/%d", srv.URL, id))
    if err != nil {
        t.Fatal(err)
    }
    
    if resp2.StatusCode != http.StatusOK {
        t.Fatalf("status %d != 200", resp2.StatusCode)
    }
}
