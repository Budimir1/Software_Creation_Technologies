# Практическое занятие №12: Подключение Swagger/OpenAPI

**ФИО:** Булискерия Будимир Бакуриевич  
**Группа:** ПИМО-01-25

Проект демонстрирует подключение Swagger/OpenAPI для автоматической генерации документации к REST API, созданному в практическом занятии №11. Реализован подход code-first с использованием аннотаций в Go-коде.

## Описание
Данный проект расширяет функциональность notes-api из ПЗ 11, добавляя автоматическую генерацию OpenAPI документации с помощью swaggo/swag. Документация публикуется через Swagger UI на эндпоинте `/docs`.

## Цели работы
1. Освоить основы спецификации OpenAPI (Swagger) для REST API.
2. Подключить автогенерацию документации к проекту notes-api.
3. Научиться публиковать интерактивную документацию (Swagger UI) на эндпоинте GET /docs.
4. Синхронизировать код и спецификацию через аннотации.
5. Подготовить процесс обновления документации.

## Требования
- Go 1.21 и выше
- Установленный swag CLI

## Установка и запуск

### 1. Установка зависимостей
```bash
# В корне проекта notes-api
go get github.com/swaggo/http-swagger
go install github.com/swaggo/swag/cmd/swag@latest

### 2. Проверка установки swag
swag -h

### 3. Добавление аннотаций в код
Добавить аннотации в следующие файлы:

cmd/api/main.go (верхнеуровневые аннотации):
// Package main Notes API server.
//
// @title Notes API
// @version 1.0
// @description Учебный REST API для заметок (CRUD).
// @contact.name Backend Course
// @contact.email example@university.ru
// @BasePath /api/v1
package main

internal/core/note.go (модели DTO):
type NoteCreate struct {
    Title   string `json:"title" example:"Новая заметка"`
    Content string `json:"content" example:"Текст заметки"`
}

type NoteUpdate struct {
    Title   *string `json:"title,omitempty" example:"Обновлено"`
    Content *string `json:"content,omitempty" example:"Новый текст"`
}

### 4. Генерация документации
swag init -g cmd/api/main.go -o docs

### 5. Подключение Swagger UI
cmd/api/main.go:
import httpSwagger "github.com/swaggo/http-swagger"

// В функции main после создания роутера:
r.Get("/docs/*", httpSwagger.WrapHandler)

### 6. Запуск сервера
go run ./cmd/api

## Использование

### Доступ к документации
Откройте в браузере: http://localhost:8080/docs/index.html

### Примеры аннотаций для обработчиков

ListNotes:
// ListNotes godoc
// @Summary Список заметок
// @Description Возвращает список заметок с пагинацией и фильтром по заголовку
// @Tags notes
// @Param page query int false "Номер страницы"
// @Param limit query int false "Размер страницы"
// @Param q query string false "Поиск по title"
// @Success 200 {array} core.Note
// @Header 200 {integer} X-Total-Count "Общее количество"
// @Failure 500 {object} map[string]string
// @Router /notes [get]

CreateNote:
// CreateNote godoc
// @Summary Создать заметку
// @Tags notes
// @Accept json
// @Produce json
// @Param input body NoteCreate true "Данные новой заметки"
// @Success 201 {object} core.Note
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /notes [post]

## Автоматизация через Makefile
Создайте файл Makefile в корне проекта:

.PHONY: run swagger

run:
    go run ./cmd/api

swagger:
    swag init -g cmd/api/main.go -o docs

Использование:
make swagger  # Генерация документации
make run      # Запуск сервера

## Структура проекта (после генерации)
notes-api/
├── cmd/api/main.go                 # Точка входа с аннотациями
├── docs/                           # Автогенерируемая документация
│   ├── docs.go                     # Инициализирующий пакет
│   ├── swagger.json                # OpenAPI спецификация (JSON)
│   └── swagger.yaml                # OpenAPI спецификация (YAML)
├── internal/
│   ├── http/
│   │   ├── router.go               # Маршрутизатор с /docs эндпоинтом
│   │   └── handlers/notes.go       # Обработчики с аннотациями
│   ├── core/
│   │   ├── note.go                 # Модели Note, NoteCreate, NoteUpdate
│   │   └── service/note_service.go # Бизнес-логика
│   └── repo/
│       └── note_mem.go             # In-memory репозиторий
├── api/openapi.yaml                # Опционально: schema-first спецификация
├── Makefile                        # Автоматизация генерации
└── go.mod