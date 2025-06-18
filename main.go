package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ImageItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	URL         string `json:"url"`
}
type Response struct {
	Data []ImageItem `json:"data"`
}

func main() {
	// Load env file theo ENV_MODE
	envFile := ".env.developer"
	if os.Getenv("ENV_MODE") == "production" {
		envFile = ".env.production"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("No .env file loaded")
	}

	// Đọc PORT từ env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "healthy", "timestamp": "` + time.Now().Format(time.RFC3339) + `"}`))
	})
	// API endpoint để lấy danh sách images
	http.HandleFunc("/api/images", func(w http.ResponseWriter, r *http.Request) {
		baseURL := os.Getenv("BASE_URL")
		response := Response{
			Data: []ImageItem{
				{
					Title:       "Learn Anywhere",
					Description: "Listen while commuting, cooking or exercising. Stay productive with your time.",
					Thumbnail:   baseURL + "/images/learn-anywhere.svg",
					URL:         baseURL + "/images/learn-anywhere.svg",
				},
				{
					Title:       "Personalized Lessons",
					Description: "Each story adapts to your level and goal with AI-driven recommendations.",
					Thumbnail:   baseURL + "/images/personalized.svg",
					URL:         baseURL + "/images/personalized.svg",
				},
				{
					Title:       "Listen& Repeat",
					Description: "10 mins activity • Now available",
					Thumbnail:   baseURL + "/images/phone.png",
					URL:         baseURL + "/images/phone.png",
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Go Image API is running", "status": "ok", "version": "1.0.0"}`))
	})
	log.Println("Listening on port " + port + "...")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
