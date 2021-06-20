package http

import "github.com/gorilla/mux"

type httpDelivery struct {
	router *mux.Router
}

func New(router *mux.Router) *httpDelivery {
	return &httpDelivery{
		router: router,
	}
}

func (hd *httpDelivery) GetToken() error {
	return nil
}
