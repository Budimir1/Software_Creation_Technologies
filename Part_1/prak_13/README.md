# Практическое занятие №13: Профилирование Go-приложения (pprof)

**ФИО:** Булискерия Будимир Бакуриевич  
**Группа:** ПИМО-01-25

Проект демонстрирует использование pprof для профилирования Go-приложений, анализ производительности функций и оптимизацию кода. Включает пример с неоптимизированной рекурсивной функцией вычисления чисел Фибоначчи и её оптимизированную итеративную версию.

## Описание
Данный проект представляет собой HTTP-сервер с подключённым pprof профилировщиком для анализа производительности. Основной фокус - демонстрация работы с различными типами профилей (CPU, Heap, Goroutine), измерение времени выполнения функций и оптимизация кода на основе полученных метрик.

## Цели работы
1. Научиться подключать и использовать профилировщик pprof для анализа CPU, памяти, блокировок и горутин.
2. Освоить базовые техники измерения времени выполнения функций.
3. Научиться читать отчёты go tool pprof, строить графы вызовов и находить "узкие места".
4. Сформировать практические навыки оптимизации кода на основании метрик.

## Требования
- Go 1.22 и выше

## Установка и запуск

### 1. Создание структуры проекта
```bash
mkdir pprof-lab
cd pprof-lab
go mod init example.com/pprof-lab
```

### 2. Структура проекта
pprof-lab/
├── cmd/api/main.go              # Точка входа с HTTP-сервером
├── internal/work/
│   ├── slow.go                  # Неоптимизированные функции
│   ├── timer.go                 # Декоратор для измерения времени
│   └── slow_test.go             # Бенчмарки
└── go.mod

### 3. Основной код сервера
cmd/api/main.go:
package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof" // Подключаем pprof
    "example.com/pprof-lab/internal/work"
)

func main() {
    mux := http.NewServeMux()
    
    // Эндпоинт с "тяжёлой" работой
    mux.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
        n := 38 // Достаточно для CPU-нагрузки
        res := work.Fib(n)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte(fmt.Sprintf("%d\n", res)))
    })

    log.Println("Server on :8080; pprof on /debug/pprof/")
    log.Fatal(http.ListenAndServe(":8080", mux))
}

### 4. Неоптимизированная функция
internal/work/slow.go:
package work

// Fib - неоптимальный рекурсивный Фибоначчи
func Fib(n int) int {
    if n < 2 {
        return n
    }
    return Fib(n-1) + Fib(n-2)
}

### 5. Запуск сервера
go run ./cmd/api

## Использование pprof

### Доступ к профилировщику
- Индекс pprof: http://localhost:8080/debug/pprof/
- CPU профиль (30 сек): http://localhost:8080/debug/pprof/profile?seconds=30
- Heap профиль: http://localhost:8080/debug/pprof/heap
- Goroutine профиль: http://localhost:8080/debug/pprof/goroutine

### Генерация нагрузки
# Установите hey (альтернатива ab)
go install github.com/rakyll/hey@latest

# Запустите нагрузку
hey -n 200 -c 8 http://localhost:8080/work
![img.png](image%2Fimg.png)

### Анализ профилей
#### Веб-интерфейс:
go tool pprof -http=:9999 http://localhost:8080/debug/pprof/profile?seconds=30

#### Консольный анализ:
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
![img_1.png](image%2Fimg_1.png)

### 1. Декоратор для измерения времени
internal/work/timer.go:
package work

import (
    "log"
    "time"
)

func TimeIt(name string) func() {
    start := time.Now()
    return func() {
        log.Printf("%s took %s", name, time.Since(start))
    }
}

Использование:
defer work.TimeIt("Fib(38)")()
res := work.Fib(38)

### 2. Бенчмарки
internal/work/slow_test.go:
package work

import "testing"

func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Fib(30)
    }
}

func BenchmarkFibFast(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = FibFast(30)
    }
}

Запуск бенчмарков:
go test -bench=. -benchmem ./...

## Оптимизация
### Оптимизированная версия функции
internal/work/slow.go (дополнение):
// FibFast - оптимизированная итеративная версия
func FibFast(n int) int {
    if n < 2 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

## Профилирование блокировок (бонус)

Включение профилирования блокировок в main.go:
import "runtime"

func init() {
    runtime.SetBlockProfileRate(1)       // Включить Block profile
    runtime.SetMutexProfileFraction(1)   // Включить Mutex profile
}

Доступ к профилям блокировок:
- http://localhost:8080/debug/pprof/block
- http://localhost:8080/debug/pprof/mutex

## Структура проекта (полная)
```
pprof-lab/
├── cmd/api/main.go                      # HTTP-сервер с pprof
├── internal/work/
│   ├── slow.go                          # Fib, FibFast
│   ├── timer.go                         # TimeIt декоратор
│   └── slow_test.go                     # Бенчмарки
├── Makefile                             # Автоматизация
├── README.md                            # Документация
└── go.mod
```

## Makefile для автоматизации
.PHONY: run profile-cpu profile-heap benchmark load-test

run:
    go run ./cmd/api

profile-cpu:
    go tool pprof -http=:9999 http://localhost:8080/debug/pprof/profile?seconds=30

profile-heap:
    go tool pprof -http=:9998 http://localhost:8080/debug/pprof/heap

benchmark:
    go test -bench=. -benchmem ./internal/work/...

load-test:
    hey -n 200 -c 8 http://localhost:8080/work

## Пример отчёта о производительности

### До оптимизации:
BenchmarkFib-8     1000    1500000 ns/op    0 B/op    0 allocs/op

### После оптимизации:
BenchmarkFibFast-8 5000000 250 ns/op        0 B/op    0 allocs/op

### Сравнительная таблица:
| Метрика        | Fib (рекурсия) | FibFast (итерация) | Ускорение |
|----------------|----------------|--------------------|-----------|
| Время (ns/op)  | 1,500,000      | 250                | 6000x     |
| Память (B/op)  | 0              | 0                  | -         |
| Аллокации      | 0              | 0                  | -         |