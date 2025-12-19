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
## Оптимизация запросов
### 1. Keyset-пагинация вместо OFFSET
internal/repo/note_pg.go:
// Плохой вариант с OFFSET (медленный на больших смещениях)
func (r *NoteRepoPostgres) ListOld(ctx context.Context, limit, offset int) ([]*core.Note, error) {
    query := `
        SELECT id, title, content, created_at
        FROM notes
        ORDER BY created_at DESC, id DESC
        LIMIT $1 OFFSET $2
    `
    // ...
}

// Оптимизированный вариант с keyset-пагинацией
func (r *NoteRepoPostgres) List(ctx context.Context, limit int, lastCreatedAt *time.Time, lastID *int64) ([]*core.Note, error) {
    var rows pgx.Rows
    var err error

    if lastCreatedAt != nil && lastID != nil {
        query := `
            SELECT id, title, content, created_at
            FROM notes
            WHERE (created_at, id) < ($1, $2)
            ORDER BY created_at DESC, id DESC
            LIMIT $3
        `
        rows, err = r.pool.Query(ctx, query, *lastCreatedAt, *lastID, limit)
    } else {
        query := `
            SELECT id, title, content, created_at
            FROM notes
            ORDER BY created_at DESC, id DESC
            LIMIT $1
        `
        rows, err = r.pool.Query(ctx, query, limit)
    }
    // ...
}

### 2. Устранение N+1 запросов (батчинг)
// До оптимизации: N отдельных запросов
func (r *NoteRepoPostgres) GetMultipleNaive(ctx context.Context, ids []int64) ([]*core.Note, error) {
    var notes []*core.Note
    for _, id := range ids {
        note, err := r.GetByID(ctx, id) // N запросов
        if err != nil {
            return nil, err
        }
        notes = append(notes, note)
    }
    return notes, nil
}

// После оптимизации: 1 запрос с батчингом
func (r *NoteRepoPostgres) GetMultiple(ctx context.Context, ids []int64) ([]*core.Note, error) {
    if len(ids) == 0 {
        return []*core.Note{}, nil
    }

    query := `
        SELECT id, title, content, created_at
        FROM notes
        WHERE id = ANY($1)
        ORDER BY created_at DESC
    `

    rows, err := r.pool.Query(ctx, query, ids)
    if err != nil {
        return nil, fmt.Errorf("get multiple notes: %w", err)
    }
    defer rows.Close()

    var notes []*core.Note
    for rows.Next() {
        var note core.Note
        if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt); err != nil {
            return nil, fmt.Errorf("scan note: %w", err)
        }
        notes = append(notes, &note)
    }

    return notes, nil
}

### 3. Подготовленные запросы (prepared statements)
type NoteRepoPostgres struct {
    pool       *pgxpool.Pool
    statements map[string]*pgconn.StatementDescription
}

func (r *NoteRepoPostgres) prepareStatements(ctx context.Context) error {
    stmts := map[string]string{
        "create_note": `
            INSERT INTO notes (title, content) 
            VALUES ($1, $2) 
            RETURNING id, created_at
        `,
        "get_note_by_id": `
            SELECT id, title, content, created_at 
            FROM notes 
            WHERE id = $1
        `,
        // ...
    }

    for name, sql := range stmts {
        sd, err := r.pool.Prepare(ctx, name, sql)
        if err != nil {
            return fmt.Errorf("prepare %s: %w", name, err)
        }
        r.statements[name] = sd
    }
    return nil
}

func (r *NoteRepoPostgres) Create(ctx context.Context, title, content string) (*core.Note, error) {
    var note core.Note
    err := r.pool.QueryRow(ctx, "create_note", title, content).
        Scan(&note.ID, &note.CreatedAt)
    if err != nil {
        return nil, fmt.Errorf("create note: %w", err)
    }
    note.Title = title
    note.Content = content
    return &note, nil
}

## Мониторинг и анализ

### 1. EXPLAIN/ANALYZE запросов
-- Анализ запроса пагинации
EXPLAIN (ANALYZE, BUFFERS)
SELECT id, title, content, created_at
FROM notes
WHERE (created_at, id) < ('2024-01-15 12:00:00', 1000)
ORDER BY created_at DESC, id DESC
LIMIT 20;

-- Анализ поиска
EXPLAIN (ANALYZE, BUFFERS)
SELECT id, title, content
FROM notes
WHERE to_tsvector('simple', title) @@ plainto_tsquery('simple', 'оптимизация');

### 2. Статистика запросов
-- Топ 10 самых тяжелых запросов
SELECT 
    query,
    calls,
    total_exec_time,
    mean_exec_time,
    rows,
    100.0 * shared_blks_hit / nullif(shared_blks_hit + shared_blks_read, 0) AS hit_percent
FROM pg_stat_statements
ORDER BY total_exec_time DESC
LIMIT 10;

### 3. Мониторинг пула подключений
internal/monitoring/pool.go:
type PoolMetrics struct {
    MaxConns     int32
    TotalConns   int32
    IdleConns    int32
    AcquiredConns int32
}

func GetPoolMetrics(pool *pgxpool.Pool) PoolMetrics {
    stats := pool.Stat()
    return PoolMetrics{
        MaxConns:     stats.MaxConns(),
        TotalConns:   stats.TotalConns(),
        IdleConns:    stats.IdleConns(),
        AcquiredConns: stats.AcquiredConns(),
    }
}

## Нагрузочное тестирование

### 1. Скрипт для тестирования
scripts/load_test.go:
package main

import (
    "fmt"
    "log"
    "sync"
    "time"

    "github.com/rakyll/hey/requester"
)

func main() {
    // Тест пагинации
    fmt.Println("Testing pagination endpoint...")
    work := &requester.Work{
        N: 1000,
        C: 50,
        Method: "GET",
        URL: "http://localhost:8080/api/v1/notes?limit=20",
    }
    work.Run()

    // Тест получения заметки
    fmt.Println("\nTesting get note endpoint...")
    work = &requester.Work{
        N: 2000,
        C: 100,
        Method: "GET",
        URL: "http://localhost:8080/api/v1/notes/1",
    }
    work.Run()
}

### 2. Makefile для автоматизации
.PHONY: migrate up down load-test monitor explain

migrate:
    psql -h localhost -U user -d notes -f migrations/001_init.sql

up:
    docker-compose up -d

down:
    docker-compose down

load-test:
    go run scripts/load_test.go

monitor:
    watch -n 2 "psql -h localhost -U user -d notes -c 'SELECT * FROM pg_stat_activity WHERE state != $$idle$$;'"

explain:
    psql -h localhost -U user -d notes -f scripts/explain_queries.sql


