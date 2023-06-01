package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type packageJSON map[string]interface{}

var fieldMap = map[string]string{
	"d":    "description",
	"dep":  "dependencies",
	"dev":  "devDependencies",
	"n":    "name",
	"peer": "peerDependencies",
	"v":    "version",
}

func main() {
	verbosePointer := flag.Bool("verbose", false, "Prints errors to the console")
	flag.Parse()

	handleError := func(err error, message string) {
		if err != nil {
			if *verbosePointer {
				fmt.Println(err)
			}
			fmt.Println(message)
			os.Exit(1)
		}
	}

	byteArray, err := os.ReadFile("package.json")
	handleError(err, "No package.json file found in this directory")

	contents := packageJSON{}
	handleError(json.Unmarshal(byteArray, &contents), "Invalid package.json file")

	var field string
	args := flag.Args()

	if len(args) > 0 {
		field = args[0]
	} else {
		field = "scripts"
	}

	if fieldMap[field] != "" {
		field = fieldMap[field]
	}

	value := contents[field]

	out, err := json.MarshalIndent(value, "", "  ")
	handleError(err, "Unable to format JSON")

	fmt.Println(string(out))
}
