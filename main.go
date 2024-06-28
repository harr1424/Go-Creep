package main

import (
	"log"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

type UserData struct {
	UserAgent string `json:"userAgent"`
	Screen    struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"screen"`
	Language string `json:"language"`
	Timezone string `json:"timezone"`
	Referrer string `json:"referrer"`
	Date     string `json:"date"`
}

type IPData struct {
	IP        string  `json:"ip"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type FullData struct {
	UserData `json:"userData"`
	IPData   `json:"ipData"`
}

var rateLimiter = NewRateLimiter()

var homePageVisitors = make([]FullData, 0)
var aboutPageVisitors = make([]FullData, 0)
var academicPortfolioPageVisitors = make([]FullData, 0)
var EULAPageVisitors = make([]FullData, 0)
var blogPrivacyPageVisitors = make([]FullData, 0)
var weatherPrivacyPageVisitors = make([]FullData, 0)

func main() {
	http.HandleFunc("/HomePage", receiveDataFromHomePage)
	http.HandleFunc("/AboutPage", receiveDataFromHAboutPage)
	http.HandleFunc("/AcademicPage", receiveDataFromAcademicPortfolioPage)
	http.HandleFunc("/EULAPage", receiveDataFromEULAPage)
	http.HandleFunc("/BlogPrivacyPage", receiveDataFromBlogPrivacyPage)
	http.HandleFunc("/WeatherPrivacyPage", receiveDataFromWeatherPrivacyPage)

	http.HandleFunc("/GetUserInfo", rateLimited(getData))

	log.Println("Server is listening on port 4141...")
	if err := http.ListenAndServe(":4141", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
