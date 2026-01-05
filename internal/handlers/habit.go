package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/blueberry-adii/momentum/internal/models"
)

type Response struct {
	Status  int
	Data    any
	Message string
	Success bool
}

func GetAllHabits(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// TODO: Get all habits from Database
	var habits []models.Habit

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Data:    habits,
		Message: "Fetched all the habits successfully",
		Success: true,
	})
}
