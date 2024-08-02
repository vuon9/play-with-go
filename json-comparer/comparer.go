package jsoncomparer

import (
	"errors"
	"fmt"
	"reflect"
)

func Compare(expected, actual interface{}) error {
	switch exp := expected.(type) {
	case map[string]interface{}:
		act, ok := actual.(map[string]interface{})
		if !ok {
			return errors.New("actual is not an object as expected")
		}

		for key, expVal := range exp {
			actVal, ok := act[key]
			if !ok {
				return fmt.Errorf("actual is missing key `%s` as expected", key)
			}

			if err := Compare(expVal, actVal); err != nil {
				return updateOrNewCompareError(err, key)
			}
		}
	case []interface{}:
		act, ok := actual.([]interface{})
		if !ok {
			return errors.New("actual is not an array as expected")
		}

		for index, expVal := range exp {
			if len(act) < index+1 {
				return fmt.Errorf("actual has no item at index %d as expected", index)
			}

			actVal := act[index]
			if err := Compare(expVal, actVal); err != nil {
				return updateOrNewCompareError(err, fmt.Sprintf("[%d]", index))
			}
		}
	default:
		// could be a string?
		if !reflect.DeepEqual(expected, actual) {
			// if not a map or slice, will go here
			return fmt.Errorf("actual is different to expected\nexpected: %+v\nactual: %+v", expected, actual)
		}
	}

	return nil
}
