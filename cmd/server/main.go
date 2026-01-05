package main

import (
	"net/http"

	"github.com/blueberry-adii/momentum/internal/database"
)

func main() {
	mux := http.NewServeMux()
	database.Connect()

	http.ListenAndServe(":8080", mux)
}
