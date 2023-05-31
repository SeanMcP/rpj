package main

import (
	"encoding/json"
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
	var isVerbose bool
	for _, arg := range os.Args {
		if arg == "--verbose" {
			isVerbose = true
		}
	}

	handleError := func(err error, message string) {
		if err != nil {
			if isVerbose {
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

	if len(os.Args) > 1 {
		field = os.Args[1]
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
