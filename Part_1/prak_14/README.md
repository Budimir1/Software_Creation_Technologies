# Практическое занятие №14: Оптимизация запросов к БД. Использование connection pool

**ФИО:** Булискерия Будимир Бакуриевич  
**Группа:** ПИМО-01-25

Проект демонстрирует оптимизацию запросов к PostgreSQL в Go-приложении, настройку пула подключений и применение различных техник для повышения производительности базы данных.

## Описание
Данный проект расширяет функциональность notes-api, добавляя работу с PostgreSQL вместо in-memory хранилища. Основное внимание уделено оптимизации SQL-запросов, настройке connection pool, использованию индексов, keyset-пагинации и устранению N+1 проблем.

## Цели работы
1. Научиться находить «узкие места» в SQL-запросах и устранять их.
2. Освоить настройку пула подключений (connection pool) в Go.
3. Научиться использовать EXPLAIN/ANALYZE и базовые метрики.
4. Применить техники уменьшения N+1 запросов и сокращения аллокаций.

## Стек технологий
- Go 1.22+
- PostgreSQL 16
- pgx/pgxpool драйвер
- Docker Compose
- hey/wrk для нагрузочного тестирования

## Установка и запуск

### 

Запуск проекта:
Клоним репозиторий:
git clone https://github.com/Budimir1/Software_Creation_Technologies/tree/main/prak_14
Проверяем что Go и Git есть:
pish_golang % cd prak_fourteen
prak_fourteen % go version
go version go1.23.2 darwin/arm64
prak_fourteen % git --version
git version 2.39.5 (Apple Git-154)
prak_fourteen %
Переходим в девятую работу:
cd prak_fourteen/cmd
Пример нашего .env:
# Remote Postgres
DB_DSN=postgres://root:root@http://address:5432/pz9_bcrypt?sslmode=disable

# HTTP
HTTP_ADDR=:8087

# Redis (локально через docker compose)
REDIS_ADDR=127.0.0.1:6379
REDIS_PASSWORD=kek
REDIS_DB=0
CACHE_TTL_SECONDS=45
Запуск на удалённом сервере:
go build -o bin/prak_fourteen ./cmd/api/

nohup ./bin/prak_fourteen >prak_fourteen.log 2>&1 &
проверка что всё равботает (пикча 3):
curl -s http://178.72.139.210:8087/health
![img.png](image%2Fimg.png)
![img_1.png](image%2Fimg_1.png)
## Создадим заметку
![img_1.png](image%2Fimg_2.png)
## Получить по Id, где второй должен кэшироваться
![img_2.png](image%2Fimg_3.png)
## Список (keyset) + поиск по title (FTS)
![img_3.png](image%2Fimg_4.png)
## Список (keyset) + поиск по title (FTS)
![img_4.png](image%2Fimg_5.png)
## Батч вместо N+1
![img_5.png](image%2Fimg_6.png)
## 


