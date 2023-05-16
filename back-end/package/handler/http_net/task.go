package http_net

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func (h *Handler) task(w http.ResponseWriter, r *http.Request) {	
	idTest := r.URL.Query().Get("idTest")
	idUser := r.URL.Query().Get("idUser")
	answer := r.URL.Query().Get("answer")
	corrAnswer := r.URL.Query().Get("corrAnswer")

	err := h.service.AddAnswer(idTest, idUser, answer, corrAnswer)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	idTask := r.URL.Query().Get("idTask")
	idVariant := r.URL.Query().Get("idVariant")
	
	task, err := h.service.Task(idUser, idTask, idVariant, idTest)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	t, err := json.Marshal(task)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Write(t)
}