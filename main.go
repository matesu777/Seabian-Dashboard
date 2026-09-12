package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matesu777/Seabian-dashboard/internal/web/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("GET /", handlers.Dashboard)

	fmt.Printf("Server running in http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
