# go-parse

In the project it was described subjects like :

- 1. How to do Http GET request to external url
- 2. Convert Json to list objects
- 3. Convert objects to JSON / XML / Text view 

## Install

 go get github.com/andriylesch/go-parse


## Http Get request

As example : This application is using external data, which is loading from

Request : 
```http
GET http://jsonplaceholder.typicode.com/photos
```

Response :
```json
200 OK
[
  {
    "albumId": 1,
    "id": 1,
    "title": "accusamus beatae ad facilis cum similique qui sunt",
    "url": "http://placehold.it/600/92c952",
    "thumbnailUrl": "http://placehold.it/150/30ac17"
  },
  ...
]

```

## Convert Json to list objects

For converting Json to list of objects was created model : 

```go

type Photo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

``` 

Part of code with converting Json to list of objects

```go

func GetPhotos(url string) ([]models.Photo, error) {

	resp, err := http.Get(url)
	
	... 

	var result []models.Photo

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return result, nil
}

```

## Convert objects to JSON / XML / Text views

### Convert to JSON 

Code :
```go
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
``` 

Json as result :
```json
{
    "total_count": 3,
    "items": [
        {
            "id": 1,
            "title": "accusamus beatae ad facilis cum similique qui sunt",
            "links": [
                {
                    "rel": "url",
                    "href": "http://placehold.it/600/92c952"
                },
                {
                    "rel": "thumbnailurl",
                    "href": "http://placehold.it/150/30ac17"
                }
            ]
        }
	
	... 
    ]
}
``` 

### Convert to XML 

Code :
```go
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
``` 

XML as result :
```xml
<PhotosXml>
    <totalcount>3</totalcount>
    <items>
        <id>1</id>
        <title>accusamus beatae ad facilis cum similique qui sunt</title>
        <url>http://placehold.it/600/92c952</url>
        <thumbnailUrl>http://placehold.it/150/30ac17</thumbnailUrl>
    </items>
</PhotosXml>
``` 

### Convert to TEXT 

Code :
```go
func totalCount(list []models.Photo) int {
	return len(list)
}

func ConvertToTextAndPrint(list []models.Photo) {

	texttemplate := template.Must(template.New("list").
		Funcs(template.FuncMap{"TotalCount": totalCount}).
		Parse(`
	        Total count : {{. | TotalCount }}
            {{range .}}----------------------------------------
	            Id : {{.Id}}
	            Title : {{.Title}}
	            Picture Url : {{.Url}}
            {{end}}`))

	if err := texttemplate.Execute(os.Stdout, list); err != nil {
		fmt.Println("Fatal error :", err)
	}
}
``` 

TEXT as result :
```text
                Total count : 2
            ----------------------------------------
                    Id : 1
                    Title : accusamus beatae ad facilis cum similique qui sunt
                    Picture Url : http://placehold.it/600/92c952
            ----------------------------------------
                    Id : 2
                    Title : reprehenderit est deserunt velit ipsam
                    Picture Url : http://placehold.it/600/771796
``` 
  
