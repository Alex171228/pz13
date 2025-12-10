package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof" // регистрирует /debug/pprof/*
    "time"

    "example.com/pprof-lab/internal/work"
)

func main() {
    mux := http.NewServeMux()

    // Эндпоинт, вызывающий "тяжёлую" работу (медленный Fibonacci).
    mux.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
        n := 38 // по заданию достаточно тяжело для CPU

        // Ручное измерение времени выполнения запроса.
        defer work.TimeIt(fmt.Sprintf("Fib(%d)", n))()

        res := work.Fib(n)

        w.Header().Set("Content-Type", "text/plain")
        _, _ = w.Write([]byte(fmtInt(res)))
    })

    // Дополнительно можно включить профили блокировок/мьютексов.
    // work.EnableLocksProfiling()

    log.Println("Server on :8080; pprof on /debug/pprof/")
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func fmtInt(v int) string { return fmt.Sprintf("%d
", v) }

// чтобы импорт time не считался неиспользованным в случае изменений
var _ = time.Now
