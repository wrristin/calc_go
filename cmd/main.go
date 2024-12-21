package main

import (
	"calc_service/internal/application"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", application.CalculateHandler)
	http.ListenAndServe(":8080", nil)
}
