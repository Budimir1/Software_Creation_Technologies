package main

import (
	"context"
	"database/sql"
	"time"
)

// Task — модель для сканирования результатов SELECT
type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo { return &Repo{DB: db} }

// CreateTask — параметризованный INSERT с возвратом id
func (r *Repo) CreateTask(ctx context.Context, title string) (int, error) {
	var id int
	const q = `INSERT INTO tasks (title) VALUES ($1) RETURNING id;`
	err := r.DB.QueryRowContext(ctx, q, title).Scan(&id)
	return id, err
}

// ListTasks — базовый SELECT всех задач (демо для занятия)
func (r *Repo) ListTasks(ctx context.Context) ([]Task, error) {
	const q = `SELECT id, title, done, created_at FROM tasks ORDER BY id;`
	rows, err := r.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *Repo) ListDone(ctx context.Context, done bool) ([]Task, error) {
	const q = `SELECT id, title, done, created_at FROM tasks WHERE done = $1 ORDER BY id`
	rows, err := r.DB.QueryContext(ctx, q, done)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// FindByID возвращает задачу по указанному ID
func (r *Repo) FindByID(ctx context.Context, id int) (*Task, error) {
	const q = `SELECT id, title, done, created_at FROM tasks WHERE id = $1`

	var task Task
	err := r.DB.QueryRowContext(ctx, q, id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// CreateMany выполняет массовую вставку задач через транзакцию
func (r *Repo) CreateMany(ctx context.Context, titles []string) error {
	// Начинаем транзакцию
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Гарантируем откат транзакции при ошибке
	defer tx.Rollback()

	// Подготавливаем запрос INSERT
	const query = `INSERT INTO tasks (title) VALUES ($1)`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполняем вставку для каждого заголовка
	for _, title := range titles {
		_, err := stmt.ExecContext(ctx, title)
		if err != nil {
			return err
		}
	}

	// Фиксируем транзакцию
	return tx.Commit()
}
