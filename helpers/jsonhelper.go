package helpers

import (
	"encoding/json"
	"fmt"
	"go-parse/models"
)

// Json models
type PhotosJson struct {
	TotalCount int         `json:"total_count"`
	Items      []PhotoJson `json:"items"`
}

type PhotoJson struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Links []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func (photo *PhotoJson) add(typeOfRel, url string) {

	if len(url) > 0 {
		var linkObj Link
		linkObj.Rel = typeOfRel
		linkObj.Href = url
		photo.Links = append(photo.Links, linkObj)
	}
}

func convertToPhotosJson(list []models.Photo) PhotosJson {

	var result PhotosJson

	if count := len(list); count > 0 {
		result.TotalCount = count
	}

	for _, item := range list {

		photoObj := PhotoJson{
			Id:    item.Id,
			Title: item.Title,
		}

		photoObj.add("url", item.Url)
		photoObj.add("thumbnailurl", item.ThumbnailUrl)

		result.Items = append(result.Items, photoObj)
	}

	return result
}

func ConvertToJSONAndPrint(list []models.Photo) {

	photosObj := convertToPhotosJson(list)

	result, _ := json.MarshalIndent(photosObj, "", "    ")

	fmt.Println(string(result))
}
