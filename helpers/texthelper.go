package helpers

import (
	"fmt"
	"go-parse/models"
	"os"
	"text/template"
)

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
