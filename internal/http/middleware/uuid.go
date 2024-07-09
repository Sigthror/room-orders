package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

var (
	uuidCtxKey = uuidKeyStruct{}
)

type uuidKeyStruct struct{}

func UUID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		ctx := context.WithValue(r.Context(), uuidCtxKey, uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
