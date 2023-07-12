package api

import "net/http"

// 账户
type User struct {
	name string
	pass string
}

// auth failed
func AuthFailed(w http.ResponseWriter, errMsg string) {
	w.Header().Set("WWW-Authenticate", `Basic realm="My REALM"`)
	w.WriteHeader(401)
	w.Write([]byte(errMsg))
}

// Basic Auth
func BasicAuthValidateMiddleware(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Basic Auth
		user, pass, ok := r.BasicAuth()
		if !ok {
			AuthFailed(w, "401 Unauthorized!")
			return
		}
		// set user
		users := make(map[string]string)
		users["admin"] = "4321"
		sysPass, exist := users[user]
		if !exist || pass != sysPass {
			AuthFailed(w, "401 Bad Password!")
			return
		}
		h(w, r)
	}
}
