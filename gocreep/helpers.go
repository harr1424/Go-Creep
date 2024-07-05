package GoCreep

import (
	"encoding/json"
	"io"
	"net/http"
)


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
