package handlers

import (
	"encoding/json"
	"ex01/models"
	"net/http"
	"strconv"
)

var candyPrices = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 5,
	"DE": 20,
	"YR": 23,
}

type CandyRequest struct {
    CandyType string `json:"candy_type"`
    Count     int    `json:"count"`
    Money     int    `json:"money"`
}

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	var req CandyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	if req.Count <= 0 {
		http.Error(w, `{"error": "Invalid candy count"}`, http.StatusBadRequest)
		return
	}

	price, validType := candyPrices[req.CandyType]
	if !validType {
		http.Error(w, `{"error": "Invalid candy type"}`, http.StatusBadRequest)
		return
	}

	totalCost := price * req.Count
	if req.Money < totalCost {
		neededMoney := totalCost - req.Money
		errorMsg := `{"error": "You need ` + strconv.Itoa(neededMoney) + ` more money!"}`
		http.Error(w, errorMsg, http.StatusPaymentRequired)
		return
	}

	change := req.Money - totalCost
	response := models.ThanksAndChange{
		Thanks: "Thank you!",
		Change: change,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
