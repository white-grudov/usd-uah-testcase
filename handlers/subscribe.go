package handlers

import (
	"encoding/json"
	"net/http"
	"usd-uah-testcase/db"
	"usd-uah-testcase/models"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	subscriber := models.Subscriber{Email: email}
	err := db.AddSubscriber(subscriber)
	if err != nil {
		if err.Error() == "unique_violation" {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Error adding subscriber", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("E-mail added")
}
