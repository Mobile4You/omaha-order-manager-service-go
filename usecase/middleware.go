package usecase

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func ensureHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if number := r.Header.Get("logic_number"); len(strings.TrimSpace(number)) == 0 {
			respondWithError(w, http.StatusBadRequest, "logic_number not found")
			log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
			return
		}

		if merchantID := r.Header.Get("merchant_id"); len(strings.TrimSpace(merchantID)) == 0 {
			respondWithError(w, http.StatusBadRequest, "merchant_id not found")
			log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
			return
		}
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)

		log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
