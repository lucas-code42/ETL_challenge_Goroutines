package main

import (
	"fmt"
	"log"
	"net/http"
	"random-numbers-api/config"
	"random-numbers-api/routers"
)

func main() {
	config.LoadConfig()
	fmt.Printf("Initialing numbers_api at the port -> %s\n", config.APIPort)
	r := routers.LoadRouters()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.APIPort), r))
}
