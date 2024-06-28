package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

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

func writeVisitorDatatoFiles() error {
	if err := writeDataToFile("homePageVisitors.json", homePageVisitors); err != nil {
		return fmt.Errorf("error writing homePageVisitors: %v", err)
	}

	if err := writeDataToFile("aboutPageVisitors.json", aboutPageVisitors); err != nil {
		return fmt.Errorf("eror writing aboutPageVisitors: %v", err)
	}

	if err := writeDataToFile("academicPortfolioPageVisitors.json", academicPortfolioPageVisitors); err != nil {
		return fmt.Errorf("error writing academicPortfolioPageVisitors: %v", err)
	}

	if err := writeDataToFile("EULAPageVisitors.json", EULAPageVisitors); err != nil {
		return fmt.Errorf("error writing EULAPageVisitors: %v", err)
	}

	if err := writeDataToFile("blogPrivacyPageVisitors.json", blogPrivacyPageVisitors); err != nil {
		return fmt.Errorf("error writing blogPrivacyPageVisitors: %v", err)
	}

	if err := writeDataToFile("weatherPrivacyPageVisitors.json", weatherPrivacyPageVisitors); err != nil {
		return fmt.Errorf("error writing weatherPrivacyPageVisitors: %v", err)
	}

	return nil
}