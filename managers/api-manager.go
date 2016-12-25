package managers

import (
	"encoding/json"
	"fmt"
	"go-parse/models"
	"net/http"
)

func GetPhotos(url string) ([]models.Photo, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request was failed: %s", resp.Status)
	}

	var result []models.Photo

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return result, nil
}
