package router

import (
	"etl-api/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRouters contem a rota e a função associada a ela
func HandleRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/sort-numbers", controller.Answer).Methods(http.MethodGet)
	return router
}
