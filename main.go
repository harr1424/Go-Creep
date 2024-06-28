package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type UserData struct {
	UserAgent string `json:"userAgent"`
	Screen    struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"screen"`
	Language string `json:"language"`
	Timezone string `json:"timezone"`
	Referrer string `json:"referrer"`
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

var homePageVisitors = make([]FullData, 0)
var aboutPageVisitors = make([]FullData, 0)
var academicPortfolioPageVisitors = make([]FullData, 0)
var EULAPageVisitors = make([]FullData, 0)
var blogPrivacyPageVisitors = make([]FullData, 0)
var weatherPrivacyPageVisitors = make([]FullData, 0)

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

func writeVisitorDatatoFiles() error {
	if err := writeDataToFile("homePageVisitors.json", homePageVisitors); err != nil {
		return fmt.Errorf("Error writing homePageVisitors: %v", err)
	}

	if err := writeDataToFile("aboutPageVisitors.json", aboutPageVisitors); err != nil {
		return fmt.Errorf("Error writing aboutPageVisitors: %v", err)
	}

	if err := writeDataToFile("academicPortfolioPageVisitors.json", academicPortfolioPageVisitors); err != nil {
		return fmt.Errorf("Error writing academicPortfolioPageVisitors: %v", err)
	}

	if err := writeDataToFile("EULAPageVisitors.json", EULAPageVisitors); err != nil {
		return fmt.Errorf("Error writing EULAPageVisitors: %v", err)
	}

	if err := writeDataToFile("blogPrivacyPageVisitors.json", blogPrivacyPageVisitors); err != nil {
		return fmt.Errorf("Error writing blogPrivacyPageVisitors: %v", err)
	}

	if err := writeDataToFile("weatherPrivacyPageVisitors.json", weatherPrivacyPageVisitors); err != nil {
		return fmt.Errorf("Error writing weatherPrivacyPageVisitors: %v", err)
	}

	return nil
}

func writeDataToFile(filename string, data []FullData) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // For pretty printing
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("could not encode data to JSON: %v", err)
	}

	return nil
}

func createZipArchive(files []string, output string) error {
	zipfile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("could not create zip file: %v", err)
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	for _, file := range files {
		err := addFileToZip(zipWriter, file)
		if err != nil {
			return fmt.Errorf("could not add file to zip: %v", err)
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	w, err := zipWriter.Create(filepath.Base(filename))
	if err != nil {
		return fmt.Errorf("could not create zip writer: %v", err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		return fmt.Errorf("could not copy file content: %v", err)
	}

	return nil
}

func main() {
	http.HandleFunc("/HomePage", receiveDataFromHomePage)
	http.HandleFunc("/AboutPage", receiveDataFromHAboutPage)
	http.HandleFunc("/AcademicPage", receiveDataFromAcademicPortfolioPage)
	http.HandleFunc("/EULAPage", receiveDataFromEULAPage)
	http.HandleFunc("/BlogPrivacyPage", receiveDataFromBlogPrivacyPage)
	http.HandleFunc("/WeatherPrivacyPage", receiveDataFromWeatherPrivacyPage)

	http.HandleFunc("/GetUserInfo", getData)

	http.ListenAndServe(":4141", nil)
}

