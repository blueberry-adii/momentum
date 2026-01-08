package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/blueberry-adii/momentum/internal/enums/frequency"
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

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		UserID      int
		Name        string
		Description string
		Frequency   frequency.Frequency
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		json.NewEncoder(w).Encode(Response{Status: http.StatusBadRequest, Data: err, Message: err.Error(), Success: false})
		return
	}

	// TODO: Create a habit into database

	w.WriteHeader(http.StatusOK)
}

func GetHabitByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	if len(path) < 3 {
		http.Error(w, "Invalid Path", http.StatusNotFound)
		return
	}

	HabitID := path[2]

	// TODO: Get Habit by ID from Database

	var habit models.Habit

	json.NewEncoder(w).Encode(Response{Status: http.StatusOK, Data: habit, Message: "Fetched Habit successfully", Success: true})
}

func DeleteHabitByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	if len(path) < 3 {
		http.Error(w, "Invalid Path", http.StatusNotFound)
		return
	}

	HabitID := path[2]

	// TODO: Delete Habit in the database

	w.WriteHeader(http.StatusOK)
}
