package middleware

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received Request: [%s] to path: %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
