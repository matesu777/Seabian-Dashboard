package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matesu777/Seabian-dashboard/web/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fs := http.FileServer(http.Dir("./web/static/"))

	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("GET /", handlers.Dashboard)
	http.HandleFunc("GET /system-stats", handlers.SystemStats)
	http.HandleFunc("GET /storage-stats", handlers.StorageStats)
	http.HandleFunc("GET /docker-services", handlers.DockerServices)
	http.HandleFunc("GET /top", handlers.TopBar)

	log.Printf("Server running in http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
