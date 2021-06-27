package http

import (
	"encoding/json"
	"github.com/AnyKey/sub-service/user"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type httpDelivery struct {
	usecase user.Usecase
}

func Launch(router *mux.Router, uc user.Usecase) error {
	userHandler := &httpDelivery{usecase: uc}
	router.HandleFunc("/token", userHandler.GetToken).Methods(http.MethodGet, http.MethodOptions)
	return nil
}
func (hd *httpDelivery) GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "token")
	token := r.Header.Get("token")
	res, err := hd.usecase.Token(token)
	if err != nil {
		log.Errorln("[GetToken] Error: ", err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}
	bytes, _ := json.Marshal(res)
	w.Write(bytes)
	return
}
