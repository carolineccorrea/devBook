package router

import "github.com/gorilla/mux"

func GerarRotas() *mux.Router {
	return mux.NewRouter()
}
