package main

import (
	"fmt"
	"strconv"
	"strings"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

type MapKey struct {
	srcKey string
	desKey string
}

type jsonMapper struct{}

func NewJSONMapper() *jsonMapper {
	return &jsonMapper{}
}

func (j *jsonMapper) ConvertToDestKeymap(srcContent string, jsonKeyMaps []MapKey) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	for _, build := range jsonKeyMaps {
		nestedKeys := strings.Split(build.desKey, ".")
		closestParent := out
		var latestItem interface{}

		for i, v := range nestedKeys {
			jq := gojsonq.New().FromString(srcContent)
			_, exists := closestParent[v]
			if exists {
				closestParent = out
				continue
			}

			if i < len(nestedKeys)-1 {
				closestParent[v] = make(map[string]interface{})
				closestParent = closestParent[v].(map[string]interface{})
				continue
			}

			item := jq.Find(build.srcKey)
			switch expectedItem := item.(type) {
			case string, int, float64, bool:
				// assign string of src item to des item
				latestItem = expectedItem
			case []interface{}:
				n, found := takeSuffixNumberIfExists(v)
				if found {
					// the des slice index is out of range
					if len(expectedItem) <= n {
						return nil, fmt.Errorf("couldn't find the index of key: %v", v)
					}

					// assign specific index of src item to des item
					latestItem = expectedItem[n]
				} else {
					// assign whole src item to des item
					latestItem = expectedItem
				}
			case nil:
				return nil, fmt.Errorf("couldn't find the key: %v", v)
			}

			closestParent[v] = latestItem
		}

	}

	return out, nil
}

func takeSuffixNumberIfExists(v string) (int, bool) {
	potentialNumber := ""
	// if the suffix is numeric character, so take it as the index could be use to query in source data
	// for example: book1 -> book[1]
	isDigit := func(c byte) bool { return c >= 0 && c <= 9 }
	for i := len(v) - 1; i > 0; i-- {
		if i != len(v)-1 && !isDigit(v[i]) {
			break
		}
		if v[i] >= 0 && v[i] <= 9 {
			potentialNumber = potentialNumber + string(v[i])
		}
	}

	if potentialNumber == "" {
		return 0, false
	}

	n64, _ := strconv.ParseInt(potentialNumber, 10, 64)

	return int(n64), true
}
