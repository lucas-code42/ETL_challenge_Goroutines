package controller

import (
	"math/rand"
	"net/http"
	"random-numbers-api/answers"
)

func GerenateNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var listNumbers []int
	for i := 0; i < 10000000; i++ {
		listNumbers = append(listNumbers, rand.Int())
	}
	answers.JSON(w, http.StatusCreated, listNumbers)
}
