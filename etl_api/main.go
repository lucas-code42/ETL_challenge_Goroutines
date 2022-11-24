package main

import (
	"etl-api/config"
	"etl-api/router"
	"fmt"
	"log"
	"net/http"
)

// main inicia o serviço
func main() {
	config.LoadConfig()
	fmt.Printf("api running at -> %s port\n", config.API_PORT)
	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.API_PORT), r))
}
