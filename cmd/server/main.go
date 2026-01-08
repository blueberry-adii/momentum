package main

import (
	"log"
	"net/http"

	"github.com/blueberry-adii/momentum/internal/database"
	"github.com/blueberry-adii/momentum/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()
	database.Connect()

	mux.HandleFunc("/habits", handlers.GetAllHabits)
	mux.HandleFunc("/habit/", handlers.GetHabitByID)
	mux.HandleFunc("/habit", handlers.CreateHabit)
	mux.HandleFunc("/habit/", handlers.DeleteHabitByID)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
