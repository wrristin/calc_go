package application

import (
	"calc_service/internal/calculation"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		if err.Error() == "недопустимый символ" || err.Error() == "недопустимое выражение" {
			http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		} else {
			log.Println("Internal server error:", err)
			http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		}
		return
	}

	response := Response{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
