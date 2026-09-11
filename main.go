package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matesu777/Seabian-dashboard/internal/docker"
	"github.com/matesu777/Seabian-dashboard/internal/models"
	"github.com/matesu777/Seabian-dashboard/internal/web/components"
)

func main() {
	loadEnv()
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("/", dashboardHandle)

	log.Printf("Server running in http://localhost%s\n", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func dashboardHandle(w http.ResponseWriter, r *http.Request) {
	data, err := getMetrics(os.Getenv("API_METRICS_URL"))
	if err != nil {
		http.Error(w, "Error, metrics not found", http.StatusInternalServerError)
	}

	servicesDocker, err := docker.ListServices()
	if err != nil {
		http.Error(w, "Error when list Docker Service", http.StatusInternalServerError)
	}
	components.Dashboard(data, servicesDocker).Render(r.Context(), w)
}

func getMetrics(Url string) (models.Response, error) {
	resp, err := http.Get(Url)
	if err != nil {
		return models.Response{}, err
	}
	defer resp.Body.Close()

	data := models.Response{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return models.Response{}, err
	}

	return data, nil
}
