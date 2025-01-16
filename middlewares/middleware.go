package middlewares

import (
	"ecommerce/utils"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONResponse(w, http.StatusUnauthorized, false, "Missing token", nil)
			return
		}

		// Ensure the token is in the "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.JSONResponse(w, http.StatusUnauthorized, false, "Invalid token format", nil)
			return
		}

		tokenString := parts[1]

		// Validate the token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.JSONResponse(w, http.StatusUnauthorized, false, "Invalid or expired token", nil)
			return
		}

		// Attach claims to the request context for downstream handlers
		ctx := utils.AddClaimsToContext(r.Context(), claims)
		r = r.WithContext(ctx)

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
}
