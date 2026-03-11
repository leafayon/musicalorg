package router

import "net/http"

type Router struct {
	server *http.Server
	mux    *http.ServeMux
}

func New(address string) *Router {
	mux := http.NewServeMux()

	return &Router{
		server: &http.Server{Addr: address, Handler: mux},
		mux:    mux,
	}
}

func (router *Router) Start() error {
	return router.server.ListenAndServe()
}
