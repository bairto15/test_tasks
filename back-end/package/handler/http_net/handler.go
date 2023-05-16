package http_net

import (
	"net/http"
	"test_puzzle/package/service"
)

type Handler struct{
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.task(w, r)

	case "POST":
		h.auth(w, r)

	case "PUT":
		h.out(w, r)
	}
}
