package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jmapper := NewJSONMapper()

	srcContent := `{"person":{"first":"John","last":"Smith","books":["Go in Action","Go in Action 2"],"school":{"address":"First Street","city":"New York"}}}`
	mapKeys := []MapKey{
		{"person.first", "first"},
		{"person.last", "last"},
		{"person.books.[0]", "nana.books1"},
	}

	mOut, err := jmapper.ConvertToDestKeymap(srcContent, mapKeys)
	if err != nil {
		panic(err)
	}
	outContent, _ := json.Marshal(mOut)

	fmt.Println(string(outContent))
}
