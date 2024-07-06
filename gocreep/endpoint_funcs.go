package GoCreep

import (
	"encoding/json"
	"log"
	"net/http"
)

var homePageVisitors = make([]FullData, 0)
var aboutPageVisitors = make([]FullData, 0)
var academicPortfolioPageVisitors = make([]FullData, 0)
var EULAPageVisitors = make([]FullData, 0)
var blogPrivacyPageVisitors = make([]FullData, 0)
var weatherPrivacyPageVisitors = make([]FullData, 0)

func ReceiveDataFromHomePage(w http.ResponseWriter, r *http.Request) {
	data, err := collectUserData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	homePageVisitors = append(homePageVisitors, data)

	city := data.City
	region := data.Region
	log.Printf("New visitor to home page from: %v, %v\n", city, region)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ReceiveDataFromHAboutPage(w http.ResponseWriter, r *http.Request) {
	data, err := collectUserData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	aboutPageVisitors = append(aboutPageVisitors, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ReceiveDataFromAcademicPortfolioPage(w http.ResponseWriter, r *http.Request) {
	data, err := collectUserData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	academicPortfolioPageVisitors = append(academicPortfolioPageVisitors, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ReceiveDataFromEULAPage(w http.ResponseWriter, r *http.Request) {
	data, err := collectUserData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	EULAPageVisitors = append(EULAPageVisitors, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ReceiveDataFromBlogPrivacyPage(w http.ResponseWriter, r *http.Request) {
	data, err := collectUserData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	blogPrivacyPageVisitors = append(blogPrivacyPageVisitors, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ReceiveDataFromWeatherPrivacyPage(w http.ResponseWriter, r *http.Request) {
	data, err := collectUserData(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherPrivacyPageVisitors = append(weatherPrivacyPageVisitors, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func DownloadReport(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	providedToken := queryParams.Get("token")

	if providedToken == "" || providedToken != token {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if err := writeVisitorDataToFiles(); err == nil {
		zipFilename := "visitor_data.zip"
		files := []string{
			"homePageVisitors.json",
			"aboutPageVisitors.json",
			"academicPortfolioPageVisitors.json",
			"EULAPageVisitors.json",
			"blogPrivacyPageVisitors.json",
			"weatherPrivacyPageVisitors.json",
		}

		err := createZipArchive(files, zipFilename)
		if err != nil {
			http.Error(w, "Could not create zip archive", http.StatusInternalServerError)
			log.Println("Error creating zip archive:", err)
			return
		}

		http.ServeFile(w, r, zipFilename)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
