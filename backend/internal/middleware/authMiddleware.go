package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/config"
	"github.com/AdarshJha-1/Taskify/backend/internal/response"
	"github.com/AdarshJha-1/Taskify/backend/internal/utils"
)

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token, err := r.Cookie("token")
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			res := response.Response{Status: http.StatusUnauthorized, Message: "Unauthorized", Data: map[string]interface{}{"error": "Unauthorized"}}
// 			json.NewEncoder(w).Encode(res)
// 			return
// 		}

// 		claims, err := utils.VerifyJWT(token.Value)
// 		if err != nil {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			res := response.Response{Status: http.StatusUnauthorized, Message: "Unauthorized", Data: map[string]interface{}{"error": err.Error()}}
// 			json.NewEncoder(w).Encode(res)
// 			return
// 		}

// 		userId := claims["user_id"].(string)
// 		fmt.Println("Extracted userID:", userId)

// 		ctx := context.WithValue(r.Context(), config.UserIDKey, userId)
// 		fmt.Println("Context value set:", ctx.Value(config.UserIDKey))
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			res := response.Response{Status: http.StatusUnauthorized, Message: "Unauthorized", Data: map[string]interface{}{"error": "Unauthorized"}}
			json.NewEncoder(w).Encode(res)
			return
		}

		claims, err := utils.VerifyJWT(token.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			res := response.Response{Status: http.StatusUnauthorized, Message: "Unauthorized", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}

		userId := claims["user_id"].(string)
		ctx := context.WithValue(r.Context(), config.UserIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
