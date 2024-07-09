package middleware

import (
	"fmt"
	"net/http"
	"time"
	"top-selection-test/internal/logger"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uuid := r.Context().Value(uuidCtxKey)
		l := logger.NewWithPrefix(logger.DebugLevel, fmt.Sprintf("REQUEST %s %s %s", uuid, r.Method, r.RequestURI))
		start := time.Now()
		defer func() {
			l.Debug("processed %s", time.Since(start))
		}()

		next.ServeHTTP(w, r.WithContext(logger.ToContext(r.Context(), l)))
	}

	return http.HandlerFunc(fn)
}
