package http_net

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) getResult(w http.ResponseWriter, r *http.Request) {
	idTest := r.URL.Query().Get("idTest")

	result, err := h.service.GetResult(idTest)

	t, err := json.Marshal(result)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Write(t)
}