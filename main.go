package main

import (
	"log"
	"net/http"

	GoCreep "github.com/harr1424/Go-Creep/gocreep"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/HomePage", GoCreep.ReceiveDataFromHomePage)
	mux.HandleFunc("/AboutPage", GoCreep.ReceiveDataFromHAboutPage)
	mux.HandleFunc("/AcademicPage", GoCreep.ReceiveDataFromAcademicPortfolioPage)
	mux.HandleFunc("/EULAPage", GoCreep.ReceiveDataFromEULAPage)
	mux.HandleFunc("/BlogPrivacyPage", GoCreep.ReceiveDataFromBlogPrivacyPage)
	mux.HandleFunc("/WeatherPrivacyPage", GoCreep.ReceiveDataFromWeatherPrivacyPage)
	mux.HandleFunc("/GetUserInfo", GoCreep.RateLimited(GoCreep.DownloadReport))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500", "http://localhost:5500"},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	log.Println("Server is listening on port 4141...")
	if err := http.ListenAndServe(":4141", handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
