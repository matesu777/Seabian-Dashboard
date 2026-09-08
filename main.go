package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matesu777/Seabian-dashboard/components"
	"github.com/matesu777/Seabian-dashboard/models"
)

func main() {
	loadEnv()
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("/", dashboard)

	fmt.Printf("Servidor rodando em http://localhost%s\n", PORT)

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

func dashboard(w http.ResponseWriter, r *http.Request) {
	data, err := getMetrics(os.Getenv("API_METRICS_URL"))
	if err != nil {
		http.Error(w, "Error ao buscar metricas", http.StatusInternalServerError)
	}
	components.Dashboard(data).Render(r.Context(), w)
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
