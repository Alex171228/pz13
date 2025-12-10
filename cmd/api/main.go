package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // включает /debug/pprof/*
	"time"

	"example.com/pprof-lab/internal/work"
)

func main() {
	mux := http.NewServeMux()

	// Эндпоинт, создающий нагрузку Fib(38)
	mux.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		n := 38

		// Ручной таймер
		defer work.TimeIt(fmt.Sprintf("Fib(%d)", n))()

		res := work.Fib(n)

		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(fmt.Sprintf("%d\n", res)))
	})

	// work.EnableLocksProfiling() // включить, если нужно профилировать блокировки

	log.Println("Server started on :8080 (pprof on /debug/pprof/ )")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// чтобы импорт time не считался неиспользованным
var _ = time.Now
