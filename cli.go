package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

var Schema = flag.String("schema", "", "Schema to validate with")    // 0
var Data = flag.String("data", "", "JSON file to check - user data") // 1
func init() {
	flag.StringVar(Schema, "s", "", "Schema to validate with")      // 0
	flag.StringVar(Data, "d", "", "JSON file to check - user data") // 1
}

func main() {

	flag.Parse()

	if *Schema == "" {
		fmt.Printf("Missing -s/--schema file name\n")
		os.Exit(1)
	}
	if *Data == "" {
		fmt.Printf("Missing -d/--data file name\n")
		os.Exit(1)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s/%s", path, *Schema))
	documentLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s/%s", path, *Data))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if !result.Valid() {
		fmt.Printf("%s is invalid.\nErrors:\n", *Data)
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		os.Exit(1)
	}

	fmt.Printf("The %s is valid\n", *Data)
	os.Exit(0)
}
