package usecase

import (
	"log"
	"net/http"
	"time"
)

func ensureBaseOrder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
		log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
