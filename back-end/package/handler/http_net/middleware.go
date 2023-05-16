package http_net

import (
	"net/http"
	"test_puzzle/package/logging"
	"test_puzzle/package/service"
)

type Middleware struct {
	handler http.Handler
	logger  logging.Logger
}

func NewMiddleware(handler http.Handler) *Middleware {
	logger := logging.GetLogger()

	return &Middleware{
		handler: handler,
		logger:  logger,
	}
}

func (l *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	login := r.Header.Get("Authorization")
	l.logger.Info(login)
	if !service.RedisAuthUser[login] && r.Method != "POST" && r.Method != "OPTIONS" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	l.handler.ServeHTTP(w, r)
}
