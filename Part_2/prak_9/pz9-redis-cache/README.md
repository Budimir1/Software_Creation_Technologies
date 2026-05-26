# pz9-redis-cache

Учебный проект по реализации распределённого кэша с Redis для backend-приложения на Go.

**Стратегия кэширования:** cache-aside  
**Кэшируемая сущность:** задача (`Task`) по ID  
**API:** `GET / PATCH / DELETE /v1/tasks/{id}`

## Особенности

- Ускорение повторного чтения задачи с использованием Redis.
- TTL с jitter для равномерного истечения ключей.
- Инвалидация кэша при обновлении (PATCH) и удалении (DELETE).
- Деградация без отказа: при недоступности Redis данные возвращаются из основного репозитория.
- Чёткое разделение: Redis не является источником истины.

## Структура проекта

```markdown
pz9-redis-cache/
├── cmd/server/main.go # точка входа
├── internal/
│ ├── config/config.go # конфигурация (адрес Redis, TTL, таймауты)
│ ├── task/
│ │ ├── model.go # модель Task
│ │ └── repo.go # in-memory репозиторий
│ ├── cache/
│ │ ├── redis.go # клиент Redis
│ │ ├── keys.go # построитель ключей
│ │ └── ttl.go # расчёт TTL с jitter
│ ├── service/
│ │ └── task_service.go # бизнес-логика (cache-aside, инвалидация)
│ └── httpapi/
│ └── handler.go # HTTP-обработчики
├── deploy/
│ └── redis/docker-compose.yml # (опционально) Docker Compose для Redis
├── go.mod
└── go.sum
```


## Требования

- Go 1.20+ (или более поздняя)
- Redis (любой способ запуска)

> **Важно:** проект использует библиотеку `github.com/redis/go-redis/v9`.

## Запуск Redis

### Способ 1 (рекомендуемый) — Memurai (нативная Windows‑версия Redis)

1. Скачайте и установите [Memurai Developer](https://www.memurai.com/).
2. После установки служба Redis будет работать на `localhost:6379`.

### Способ 2 — Docker Compose

```bash
cd deploy/redis
docker compose up -d
```
Способ 3 — Redis for Windows (legacy)
Скачайте архив Redis-x64-3.2.100.zip со страницы MicrosoftArchive, распакуйте и запустите redis-server.exe.

Убедитесь, что Redis доступен, выполнив (из папки проекта):

bash
redis-cli ping
Ожидаемый ответ: PONG.

Запуск приложения
Клонируйте репозиторий и перейдите в папку проекта.

Установите зависимости:

bash
go mod tidy
Запустите сервер:

bash
go run ./cmd/server
Сервер начнёт слушать порт 8082.
Если Redis недоступен, в консоли появится предупреждение, но сервер продолжит работу.

Проверка сценариев
Все примеры используют curl (PowerShell, Git Bash, WSL и т.п.).

1. Чтение задачи и наполнение кэша
   bash
# Первый запрос – попадание в БД, сохранение в кэш
curl http://localhost:8082/v1/tasks/1

# Второй запрос – ответ из кэша
curl http://localhost:8082/v1/tasks/1
Логи сервера:

text
cache miss: tasks:task:1
cache hit: tasks:task:1
2. Инвалидация при обновлении
   bash
   curl -X PATCH http://localhost:8082/v1/tasks/1 \
   -H "Content-Type: application/json" \
   -d "{\"id\":1,\"title\":\"Обновлённая\",\"description\":\"...\",\"due_date\":\"2026-01-22T00:00:00Z\"}"

# После PATCH – снова чтение
```bash
curl http://localhost:8082/v1/tasks/1
```
Ключ tasks:task:1 удаляется, следующим GET‑запросом данные снова попадают в кэш.

### 3. Инвалидация при удалении
   ```bash
   curl -X DELETE http://localhost:8082/v1/tasks/1
   curl http://localhost:8082/v1/tasks/1   # 404 Not Found
```
### 4. Деградация при недоступности Redis
   Остановите Redis (например, docker compose stop в папке deploy/redis, или остановите службу Memurai).

Выполните запрос:

```bash
curl http://localhost:8082/v1/tasks/2
```
Сервер не падает, возвращает данные из репозитория, в логах предупреждение о недоступности Redis.

### Дополнительные возможности
Кэширование списка задач (GET /v1/tasks).

Отрицательное кэширование (кэширование факта отсутствия сущности).

Разделение построителя ключей и сериализатора.

Расширенное логирование hit/miss/set/invalidate.

### Контрольные точки
Стратегия cache-aside реализована.

TTL + jitter применяются.

Инвалидация при изменении/удалении работает.

При отказе Redis сервис продолжает обслуживать запросы.

Единая схема ключей tasks:task:<id>.



