package main

import (
	"log"
	"net/http"
)

type Coupon struct {
	Code string
}

type Result struct {
	Status string
}

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	result := Result{Status: "Cupom disponivel: TONABLACK"}

	log.Println(result.Status)

}
