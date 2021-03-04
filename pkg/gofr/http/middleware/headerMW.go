package middleware

import (
	"net/http"
)

func DefaultHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-type", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")

		inner.ServeHTTP(w, r)

	})

}
