package main

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "time"
)

func createChiRouter() chi.Router {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        t := time.Now().Format(time.RFC3339)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(t))
    })
    return r
}

func main() {
    r := createChiRouter()
    s := http.Server{
        Addr:         ":8080",
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 90 * time.Second,
        IdleTimeout:  120 * time.Second,
        Handler:      r,
    }
    err := s.ListenAndServe()
    if err != nil {
        if err != http.ErrServerClosed {
            panic(err)
        }
    }
}
