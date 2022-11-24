package router

import "github.com/gorilla/mux"

// GenerateRouter gera uma rota e chama a HandleRouters 
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return HandleRouters(r)
}
