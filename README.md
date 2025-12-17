# pprof-lab — Практическое занятие №13

Мини-сервер на Go для отработки профилирования (pprof), измерения времени функций
и бенчмарков. Сервер поднимает HTTP-обработчик `/work`, который вызывает тяжёлую
функцию Fibonacci, и экспортирует профили по `/debug/pprof/*`.

## Структура проекта

```text
pprof-lab/
 ├─ cmd/
 │   └─ api/
 │       └─ main.go          # HTTP-сервер + подключение pprof
 ├─ internal/
 │   └─ work/
 │       ├─ slow.go          # Fib (медленный) и FibFast (быстрый)
 │       ├─ timer.go         # TimeIt: ручной таймер
 │       ├─ slow_test.go     # бенчмарки BenchmarkFib/BenchmarkFibFast
 │       └─ locks.go         # включение block/mutex профилей (доп.)
 └─ go.mod
```

## Запуск сервера

```bash
go run ./cmd/api
```

Сервер будет слушать порт `:8080`:

- Эндпоинт нагрузки: `http://localhost:8080/work`
- Профили pprof: `http://localhost:8080/debug/pprof/`

## Генерация нагрузки

В другом терминале можно создать нагрузку на `/work`, например при помощи `hey`:
<img width="555" height="822" alt="image" src="https://github.com/user-attachments/assets/65c93714-562b-4728-9da0-51ca4e675f5d" />

```bash
hey -n 200 -c 8 http://localhost:8080/work
```

Или скриптом на `curl` в цикле.

## Получение профилей

- Индекс pprof: `http://localhost:8080/debug/pprof/`
- CPU (30 секунд):

  ```bash
  go tool pprof -http=:9999 http://localhost:8080/debug/pprof/profile?seconds=30
  ```

- Heap:

  ```bash
  go tool pprof -http=:9998 http://localhost:8080/debug/pprof/heap
  ```

## Бенчмарки

```bash
go test -bench=. -benchmem ./...
```

Сравните `BenchmarkFib` и `BenchmarkFibFast` по метрикам:
- ns/op — время на операцию
- B/op — байт на операцию
- allocs/op — количество аллокаций на операцию
