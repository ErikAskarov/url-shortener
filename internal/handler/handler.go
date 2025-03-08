package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
	"url-shortener/internal/db"
	"url-shortener/internal/models"
)

type Handler struct {
	redis *db.RedisStorage
}

var (
	codeGenerator     *rand.Rand
	codeGeneratorOnce sync.Once
)

func initCodeGenerator() {
	codeGeneratorOnce.Do(func() {
		src := rand.NewSource(time.Now().UnixNano())
		codeGenerator = rand.New(src)
	})
}

func NewHandler(redis *db.RedisStorage) *Handler {
	return &Handler{redis: redis}
}

func generateCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)

	initCodeGenerator()

	for i := range b {
		b[i] = charset[codeGenerator.Intn(len(charset))]
	}
	return string(b)
}

func parseStringToDuration(exp string) (t time.Duration, err error) {
	parsed, err := time.ParseDuration(exp)
	if err != nil {
		return 0, fmt.Errorf("Failed parse value")
	}
	return parsed, nil
}

func (h *Handler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var req models.Url

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	code := generateCode()
	if req.Expires == "" {
		http.Error(w, "Invalid time duration", http.StatusBadRequest)
		return
	}
	parseTime, err := parseStringToDuration(req.Expires)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := h.redis.SaveUrl(r.Context(), req.Url, code, parseTime); err != nil {

	}
}
