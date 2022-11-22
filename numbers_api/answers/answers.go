package answers

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONAnswer struct {
	Numbers []int `json:"numbers"`
}

func JSON(w http.ResponseWriter, statuscode int, data []int) {
	numbers := JSONAnswer{data}
	w.WriteHeader(statuscode)
	body, err := json.Marshal(&numbers)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(body))
}
