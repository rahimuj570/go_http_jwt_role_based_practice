package middlewares

import (
	"net/http"

	"github.com/rahimuj570/go_http_jwt_role_based_practice/enums"
	jwtconfig "github.com/rahimuj570/go_http_jwt_role_based_practice/jwt_config"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearrer_token := r.Header.Get("Authorization")
		if bearrer_token == "" || len(bearrer_token) < 8 {
			http.Error(w, "No Header Found", http.StatusUnauthorized)
			return
		}
		string_token := bearrer_token[len("bearrer "):]
		data, role := jwtconfig.PurseJWT(string_token, w)
		//JUST FOR PRACTICE
		if role != enums.ADMIN && data != "" {
			http.Error(w, "No no no no", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
