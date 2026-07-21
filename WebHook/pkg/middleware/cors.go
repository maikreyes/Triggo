package middleware

import "net/http"

func (m *Middleware) Cors(w http.ResponseWriter, r *http.Request) bool {

	w.Header().Set("Access-Control-Allow-Origin", m.Config.FrontUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}

	return false

}
