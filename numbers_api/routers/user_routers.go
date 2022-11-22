package routers

import (
	"fmt"
	"net/http"
	"random-numbers-api/controller"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	URL      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
}

var user = []UserRouter{
	{
		URL:      "/numbers",
		Method:   http.MethodGet,
		Function: controller.GerenateNumbers,
	},
}

func UserRouters(r *mux.Router) *mux.Router {
	for _, router := range user {
		fmt.Println("Routers:", router)
		r.HandleFunc(router.URL, router.Function).Methods(router.Method)
	}
	return r
}
