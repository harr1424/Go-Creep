package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func receiveDataFromHomePage(w http.ResponseWriter, r *http.Request) {
	var data FullData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error reeiving visitor data from homepage:", err.Error())
		return
	}

	homePageVisitors = append(homePageVisitors, data)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"success"}`))
}

func receiveDataFromHAboutPage(w http.ResponseWriter, r *http.Request) {
	var data FullData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error reeiving visitor data from homepage:", err.Error())
		return
	}

	aboutPageVisitors = append(aboutPageVisitors, data)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"success"}`))
}

func receiveDataFromAcademicPortfolioPage(w http.ResponseWriter, r *http.Request) {
	var data FullData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error reeiving visitor data from homepage:", err.Error())
		return
	}

	academicPortfolioPageVisitors = append(academicPortfolioPageVisitors, data)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"success"}`))
}

func receiveDataFromEULAPage(w http.ResponseWriter, r *http.Request) {
	var data FullData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error reeiving visitor data from homepage:", err.Error())
		return
	}

	EULAPageVisitors = append(EULAPageVisitors, data)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"success"}`))
}

func receiveDataFromBlogPrivacyPage(w http.ResponseWriter, r *http.Request) {
	var data FullData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error reeiving visitor data from homepage:", err.Error())
		return
	}

	blogPrivacyPageVisitors = append(blogPrivacyPageVisitors, data)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"success"}`))
}

func receiveDataFromWeatherPrivacyPage(w http.ResponseWriter, r *http.Request) {
	var data FullData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error reeiving visitor data from homepage:", err.Error())
		return
	}

	weatherPrivacyPageVisitors = append(weatherPrivacyPageVisitors, data)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"success"}`))
}

func getData(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	providedToken := queryParams.Get("token")

	if providedToken == "" || providedToken != token {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if err := writeVisitorDatatoFiles(); err == nil {
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
