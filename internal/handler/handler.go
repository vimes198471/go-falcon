package handler

import (
    "encoding/json"
    "net/http"
    "sync"
    "time"
)

type Item struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Value     float64   `json:"value"`
    CreatedAt time.Time `json:"created_at"`
}

type Handler struct {
    mu     sync.RWMutex
    items  map[int]*Item
    nextID int
}

func NewHandler() *Handler {
    return &Handler{
        items:  make(map[int]*Item),
        nextID: 1,
    }
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status":     "ok",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}

func (h *Handler) Items(w http.ResponseWriter, r *http.Request) {
    h.mu.RLock()
    defer h.mu.RUnlock()

    items := make([]*Item, 0, len(h.items))
    for _, item := range h.items {
        items = append(items, item)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}