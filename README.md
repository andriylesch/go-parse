# go-parse

In the project it was described subjects like :

- 1. How to do Http GET request to external url
- 2. Convert Json to list objects
- 3. Convert objects to JSON / XML / Text view 

## Http Get request

As example : This application is using external data, which is loading from

Request : 
```http
GET http://jsonplaceholder.typicode.com/photos
```

Response :
```http
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

```http

type Photo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

``` 

Part of code with converting Json to list of objects

```http

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

  