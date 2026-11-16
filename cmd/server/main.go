package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "github.com/vimes198471/go-falcon/internal/handler"
    "github.com/vimes198471/go-falcon/internal/middleware"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mux := http.NewServeMux()
    h := handler.NewHandler()
    mux.HandleFunc("/api/health", h.Health)
    mux.HandleFunc("/api/items", h.Items)

    wrapped := middleware.Logging(middleware.Recover(mux))

    server := &http.Server{
        Addr:    ":" + port,
        Handler: wrapped,
    }

    go func() {
        log.Printf("Server starting on :%s", port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop
    log.Println("Shutting down...")
}