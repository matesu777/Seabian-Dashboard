package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matesu777/Seabian-dashboard/internal/docker"
	"github.com/matesu777/Seabian-dashboard/internal/models"
	"github.com/matesu777/Seabian-dashboard/web/components"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	data, err := GetMetrics(os.Getenv("API_METRICS_URL"))
	log.Printf("[API] Take metrics for Dashboard in %s", os.Getenv("API_METRICS_URL"))
	if err != nil {
		http.Error(w, "Error, metrics not found", http.StatusInternalServerError)
		log.Printf("[ERROR DASH] Dashboard metrics not found in %s", os.Getenv("API_METRICS_URL"))
		return
	}

	servicesDocker, err := docker.ListServices()
	if err != nil {
		http.Error(w, "Error when list Docker Service", http.StatusInternalServerError)
		log.Print("[ERROR DASH] Dashboard Docker Services, check your Docker App")
		return
	}
	components.DashboardPage(data, servicesDocker).Render(r.Context(), w)
}

func SystemStats(w http.ResponseWriter, r *http.Request) {
	data, err := GetMetrics(os.Getenv("API_METRICS_URL"))
	log.Printf("[API] Take metrics System in %s", os.Getenv("API_METRICS_URL"))
	if err != nil {
		http.Error(w, "failed to get metrics", http.StatusInternalServerError)
		return
	}

	components.SystemStats(data).Render(r.Context(), w)
}

func StorageStats(w http.ResponseWriter, r *http.Request) {
	data, err := GetMetrics(os.Getenv("API_METRICS_URL"))
	log.Printf("[API] Take metrics Storage in %s", os.Getenv("API_METRICS_URL"))
	if err != nil {
		http.Error(w, "failed to get metrics", http.StatusInternalServerError)
		return
	}

	components.StorageStats(data).Render(r.Context(), w)
}

func DockerServices(w http.ResponseWriter, r *http.Request) {
	servicesDocker, err := docker.ListServices()
	if err != nil {
		http.Error(w, "Error when list Docker Service", http.StatusInternalServerError)
		log.Print("[ERROR] When list Docker Service in service components")
		return
	}
	components.ServicesLayout(servicesDocker).Render(r.Context(), w)
}

func TopBar(w http.ResponseWriter, r *http.Request) {
	data, err := GetMetrics(os.Getenv("API_METRICS_URL"))
	log.Printf("[API] Take metrics Top bar in %s", os.Getenv("API_METRICS_URL"))
	if err != nil {
		http.Error(w, "failed to get metrics", http.StatusInternalServerError)
		return
	}

	components.Top(data.System).Render(r.Context(), w)
}

func Network(w http.ResponseWriter, r *http.Request) {
	data, err := GetMetrics(os.Getenv("API_METRICS_URL"))
	log.Printf("[API] Take metrics Network in %s", os.Getenv("API_METRICS_URL"))
	if err != nil {
		http.Error(w, "failed to get metrics", http.StatusInternalServerError)
		return
	}
	components.NetworkCard(data.Hardware.Network[1]).Render(r.Context(), w)
}

func GetMetrics(url string) (models.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Response{}, fmt.Errorf(
			"metrics API returned status %d",
			resp.StatusCode,
		)
	}

	var data models.Response

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.Response{}, err
	}

	return data, nil
}
