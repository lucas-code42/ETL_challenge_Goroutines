package controller

import (
	"encoding/json"
	"etl-api/models"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Answer ficará responsável por responder a requisição
func Answer(w http.ResponseWriter, r *http.Request) {
	process()
}

// process escopo global para o processamento
func process() {
	numbers := getNumbers()
	splitList(numbers.Number)
}

// getNumbers faz a requisição para a outra API local que fornece os dados
func getNumbers() models.NumbersList {
	response, err := http.Get("http://127.0.0.1:5000/numbers")
	if err != nil {
		log.Fatal(err)
	}

	var numberList models.NumbersList
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &numberList)
	if err != nil {
		log.Fatal(err)
	}

	return numberList
}

// splitList reparte a lista obtida
func splitList(numbers []int) {
	fmt.Println(numbers[0:3])
}
