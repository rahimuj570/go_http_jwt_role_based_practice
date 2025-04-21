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
		string_token := bearrer_token[len("bearer "):]
		// println(string_token)
		data, role := jwtconfig.PurseJWT(string_token, w)
		//JUST FOR PRACTICE
		role2 := string(role.(string))
		// println(role.(enums.Role))
		data2 := int(data.(float64))
		if role2 != string(enums.ADMIN) || data2 == 0 {
			http.Error(w, "No no no no", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
