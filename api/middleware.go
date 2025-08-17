package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
)

// key for storing user ID in request context
type contextKey string

const userContextKey contextKey = "userID"

// AuthMiddleware creates a middleware that verifies Firebase ID tokens.
func AuthMiddleware(app *firebase.App) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Authorization header is required"})
				return
			}

			idToken := strings.Replace(authHeader, "Bearer ", "", 1)
			client, err := app.Auth(context.Background())
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get Firebase Auth client"})
				return
			}

			token, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid or expired token"})
				return
			}

			// Add the user ID to the request context
			ctx := context.WithValue(r.Context(), userContextKey, token.UID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// GetUserIDFromContext retrieves the user ID from the request context.
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userContextKey).(string)
	return userID, ok
}