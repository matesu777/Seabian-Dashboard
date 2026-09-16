package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matesu777/Seabian-dashboard/web/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./web/static/"))

	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("GET /", handlers.Dashboard)
	http.HandleFunc("GET /system-stats", handlers.SystemStats)
	http.HandleFunc("GET /storage-stats", handlers.StorageStats)
	http.HandleFunc("GET /docker-services", handlers.DockerServices)
	http.HandleFunc("GET /top", handlers.TopBar)
	http.HandleFunc("GET /network-stats", handlers.Network)

	log.Printf("[API] Metrics API: %s", os.Getenv("API_METRICS_URL"))
	log.Printf("[START] Server running on :%s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
