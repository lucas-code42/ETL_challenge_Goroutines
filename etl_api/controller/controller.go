package controller

import (
	"encoding/json"
	"etl-api/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
)

// Answer ficará responsável por responder a requisição
func Answer(w http.ResponseWriter, r *http.Request) {
	process()
}

// process escopo global para o processamento
func process() {
	runtime.GOMAXPROCS(8)

	numbers := getNumbers()

	cut := len(numbers) / 10
	numbers, s1 := divideList(numbers, cut)

	cut = len(numbers) / 8
	numbers, s2 := divideList(numbers, cut)

	cut = len(numbers) / 6
	numbers, s3 := divideList(numbers, cut)

	cut = len(numbers) / 4
	numbers, s4 := divideList(numbers, cut)

	c := make(chan []int)
	go sortNUmbersLowerToBigger(s1, c)
	go sortNUmbersLowerToBigger(s2, c)
	go sortNUmbersLowerToBigger(s3, c)
	go sortNUmbersLowerToBigger(s4, c)

	sorted1 := <-c
	sorted2 := <-c
	sorted3 := <-c
	sorted4 := <-c

	var finalList []int
	finalList = append(finalList, sorted1...)
	finalList = append(finalList, sorted2...)
	finalList = append(finalList, sorted3...)
	finalList = append(finalList, sorted4...)
	finalList = append(finalList, numbers...)

	finalList = sortNumbers(finalList)
	fmt.Println(finalList)

}

// getNumbers faz a requisição para a outra API local que fornece os dados
func getNumbers() []int {
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

	return numberList.Number
}

// divideList reparte a lista obtida
func divideList(list []int, cut int) (tmp, slice []int) {
	tmp = list[cut:]
	slice = list[:cut]
	return tmp, slice
}

// sortNUmbersBiggerToLower ordena os numeros do maior para o menor
func sortNUmbersLowerToBigger(numbersList []int, c chan []int) {
	for j := 0; j < len(numbersList); j++ {
		for i := 0; i < len(numbersList); i++ {

			var temp int
			currentPosition := i
			nextPostion := i + 1

			if nextPostion == len(numbersList) {
				break
			}

			if numbersList[currentPosition] < numbersList[nextPostion] {
				temp = numbersList[currentPosition]
				numbersList[currentPosition] = numbersList[nextPostion]
				numbersList[nextPostion] = temp

			}
		}

	}
	c <- numbersList
}

// sortNumbers ordena os numeros do maior para o menor
func sortNumbers(numbersList []int) []int {
	for j := 0; j < len(numbersList); j++ {
		for i := 0; i < len(numbersList); i++ {

			var temp int
			currentPosition := i
			nextPostion := i + 1

			if nextPostion == len(numbersList) {
				break
			}

			if numbersList[currentPosition] < numbersList[nextPostion] {
				temp = numbersList[currentPosition]
				numbersList[currentPosition] = numbersList[nextPostion]
				numbersList[nextPostion] = temp

			}
		}

	}
	return numbersList
}
