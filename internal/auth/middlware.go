package auth 

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

func Middleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization")
				if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
						next.ServeHTTP(w, r)
						return
				}

				tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
				userID, err := ParseToken(tokenStr)
				if err == nil {
						ctx := context.WithValue(r.Context(), UserIDKey, userID)
						r = r.WithContext(ctx)
				}
				next.ServeHTTP(w, r)
	})

}

func GetUserID(ctx context.Context) (uint, bool) {
		id, ok := ctx.Value(UserIDKey).(uint)
		return id, ok
}
