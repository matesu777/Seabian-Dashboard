package handlers

import (
	"encoding/json"
	"github.com/matesu777/Seabian-dashboard/internal/docker"
	"github.com/matesu777/Seabian-dashboard/internal/models"
	"github.com/matesu777/Seabian-dashboard/internal/web/components"
	"net/http"
	"os"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
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
