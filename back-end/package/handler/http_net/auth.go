package http_net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type signin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Авторизация
func (h *Handler) auth(w http.ResponseWriter, r *http.Request) {
	req := signin{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	idUser, err := h.service.Auth(req.Login, req.Password)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	if idUser == "" {
		fmt.Fprint(w, "доступ запрещен")
		return
	}

	variants, err := h.service.GetVariants()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	res := map[string]interface{}{
		"variants": variants,
		"idUser": idUser,
	}

	v, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Write(v)
}
