package GoCreep

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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
	Date     string `json:"date"`
}

type FullData struct {
	UserData
	IP        string  `json:"ip"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func getIpInfo() (map[string]interface{}, error) {
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipData map[string]interface{}
	if err := json.Unmarshal(body, &ipData); err != nil {
		return nil, err
	}

	return ipData, nil
}

func collectUserData(_ http.ResponseWriter, r *http.Request) (FullData, error) {
	var userData UserData
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		return FullData{}, err
	}

	ipData, err := getIpInfo()
	if err != nil {
		return FullData{}, err
	}

	fullData := FullData{
		UserData:  userData,
		IP:        getStringFromMap(ipData, "ip"),
		City:      getStringFromMap(ipData, "city"),
		Region:    getStringFromMap(ipData, "region"),
		Country:   getStringFromMap(ipData, "country_name"),
		Latitude:  getFloat64FromMap(ipData, "latitude"),
		Longitude: getFloat64FromMap(ipData, "longitude"),
	}

	return fullData, nil
}

func getStringFromMap(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok && val != nil {
		return val.(string)
	}
	return ""
}

func getFloat64FromMap(m map[string]interface{}, key string) float64 {
	if val, ok := m[key]; ok && val != nil {
		return val.(float64)
	}
	return 0.0
}

func validateToken(token string) error {
    if len(token) == 0 {
        return errors.New("token is empty")
    }
    
    // Check if the token is alphanumeric
    for _, char := range token {
        if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
            return errors.New("token contains invalid characters")
        }
    }
    
    return nil
}
