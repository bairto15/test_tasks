package http_net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Разлогиниться
func (h *Handler) out(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.Out(req.Login)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, "Разлогирование успешно прошло")
}