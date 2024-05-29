package handlers

import (
	"encoding/json"
	"ex00/models"
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

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	if order.CandyCount <= 0 {
		http.Error(w, `{"error": "Invalid candy count"}`, http.StatusBadRequest)
		return
	}

	price, validType := candyPrices[order.CandyType]
	if !validType {
		http.Error(w, `{"error": "Invalid candy type"}`, http.StatusBadRequest)
		return
	}

	totalCost := price * order.CandyCount
	if order.Money < totalCost {
		neededMoney := totalCost - order.Money
		errorMsg := `{"error": "You need ` + strconv.Itoa(neededMoney) + ` more money!"}`
		http.Error(w, errorMsg, http.StatusPaymentRequired)
		return
	}

	change := order.Money - totalCost
	response := models.ThanksAndChange{
		Thanks: "Thank you!",
		Change: change,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
