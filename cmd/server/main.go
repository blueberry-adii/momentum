package main

import (
	"net/http"

	"github.com/blueberry-adii/momentum/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()
	database.Connect()

	http.ListenAndServe(":8080", mux)
}
