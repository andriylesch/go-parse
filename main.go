package main

import (
	"fmt"
	"go-parse/helpers"
	"go-parse/managers"
)

const url = "http://jsonplaceholder.typicode.com/photos"

func main() {

	fmt.Println("\n Does http get request \n")
	result, _ := managers.GetPhotos(url)

	fmt.Println("\n---Convert to JSON --- \n ")
	helpers.ConvertToJSONAndPrint(result[:3])

	fmt.Println("\n---Convert to XML ---\n")
	helpers.ConvertToXMLAndPrint(result[:3])

	fmt.Println("\n---Convert to TEXT ---\n")
	helpers.ConvertToTextAndPrint(result[:3])

}
