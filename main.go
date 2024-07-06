package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	GoCreep "github.com/harr1424/Go-Creep/gocreep"
	"github.com/rs/cors"
)

func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/HomePage", GoCreep.ReceiveDataFromHomePage).Methods("POST")
	r.HandleFunc("/AboutPage", GoCreep.ReceiveDataFromHAboutPage).Methods("POST")
	r.HandleFunc("/AcademicPage", GoCreep.ReceiveDataFromAcademicPortfolioPage).Methods("POST")
	r.HandleFunc("/EULAPage", GoCreep.ReceiveDataFromEULAPage).Methods("POST")
	r.HandleFunc("/BlogPrivacyPage", GoCreep.ReceiveDataFromBlogPrivacyPage).Methods("POST")
	r.HandleFunc("/WeatherPrivacyPage", GoCreep.ReceiveDataFromWeatherPrivacyPage).Methods("POST")
	r.HandleFunc("/GetUserInfo", GoCreep.RateLimited(GoCreep.DownloadReport)).Methods("GET")

	r.Use(securityHeadersMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500", "http://localhost:5500"}, // change accordingly
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:4141",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server is listening on port 4141...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
