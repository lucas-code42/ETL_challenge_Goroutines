package routers

import "github.com/gorilla/mux"

// LoadRouters load the routers of API
func LoadRouters() *mux.Router {
	r := mux.NewRouter()
	return UserRouters(r)
}
