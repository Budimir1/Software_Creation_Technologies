# Практическое занятие №5 - Работа с PostgreSQL в Go

## Описание проекта
Проект демонстрирует подключение к PostgreSQL из Go-приложения с использованием пакета `database/sql` и драйвера `pgx`. Включает выполнение базовых SQL-запросов (INSERT, SELECT) с правильной обработкой ошибок и использованием context.

## Технологии
- **Go 1.21**
- **PostgreSQL 18**
- **Драйвер**: github.com/jackc/pgx/v5

## Предварительные требования

### Установка PostgreSQL
- **Windows**: Установщик EnterpriseDB
- **macOS**: `brew install postgresql && brew services start postgresql`
- **Linux**: `sudo apt install postgresql`

### Проверка установки
```bash
psql --version
go version
```
## Установка и запуск проекта

### Создание директории проекта
```
mkdir prak_5 && cd prak_5
```
### Инициализация Go модуля
```
go mod init Budimir/prak_5
```
## Установка зависимостей
### Установка драйвера PostgreSQL для Go
```
go get github.com/jackc/pgx/v5/stdlib
```
### Установка библиотеки для работы с .env файлами (опционально)
```
go get github.com/joho/godotenv
```
## Структура проекта
```
pz5-db/
├── main.go          # Основной файл приложения
├── db.go            # Настройка подключения к БД
├── repository.go    # Логика работы с данными
├── go.mod           # Файл модуля Go
├── go.sum           # Файл зависимостей
└── .env             # Файл с переменными окружения (опционально)
```
## Настройка подключения к БД
### Использование .env файла
```
DATABASE_URL=postgres://postgres:YOUR_PASSWORD@localhost:5432/todo?sslmode=disable
```
