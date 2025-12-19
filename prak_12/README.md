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

## Установка и 

## Структура проекта (после генерации)
```
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
```

Запуск проекта:
Клоним репозиторий:
git clone 
Проверяем что Go и Git есть:
pish_golang % cd prak_twelve

prak_twelve % go version
go version go1.23.2 darwin/arm64

prak_twelve % git --version
git version 2.39.5 (Apple Git-154)

prak_twelve %
создаём SWAGGER:
swag init -g cmd/api/main.go -o docs
Стоит отметить, что есть ошибка, связанная с LeftDelim и RightDelim (cносим их к чёртовому С В А Г Г Е Р У)

Демонстрация:
Скриншот работающей страницы Swagger UI

![img.png](image%2Fimg.png)
![img_2.png](image%2Fimg_2.png)
![img_1.png](image%2Fimg_1.png)
![img_3.png](image%2Fimg_3.png)
![img_4.png](image%2Fimg_4.png)

[sponge-bob-finished.mp4](image_5%2Fsponge-bob-finished.mp4)

