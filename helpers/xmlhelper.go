package helpers

import (
	"encoding/xml"
	"fmt"
	"go-parse/models"
)

type PhotosXml struct {
	TotalCount int        `xml:"totalcount"`
	Items      []PhotoXml `xml:"items"`
}

type PhotoXml struct {
	Id           int    `xml:"id"`
	Title        string `xml:"title"`
	Url          string `xml:"url"`
	ThumbnailUrl string `xml:"thumbnailUrl"`
}

func ConvertToXMLAndPrint(list []models.Photo) {

	var photosXml PhotosXml

	photosXml.TotalCount = len(list)

	for _, item := range list {

		photoXml := PhotoXml{
			Id:           item.Id,
			Title:        item.Title,
			Url:          item.Url,
			ThumbnailUrl: item.ThumbnailUrl,
		}

		photosXml.Items = append(photosXml.Items, photoXml)
	}

	result, err := xml.MarshalIndent(photosXml, "", "    ")
	if err != nil {
		fmt.Println("Fatal error during converting to XML :", err)
		return
	}

	fmt.Println(string(result))
}
