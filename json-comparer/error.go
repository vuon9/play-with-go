package jsoncomparer

import (
	"fmt"
	"strings"
)

type CompareErr struct {
	err  error
	keys []string
}

func (e *CompareErr) PrependKey(key string) {
	e.keys = append([]string{key}, e.keys...)
}

// Error prints the key of mismatch property with details error
func (e *CompareErr) Error() string {
	return fmt.Sprintf("failed at `%s`: %s", strings.Join(e.keys, "."), e.err.Error())
}

func updateOrNewCompareError(err error, key string) *CompareErr {
	if cErr, ok := err.(*CompareErr); ok {
		cErr.PrependKey(key)
		return cErr
	}

	return &CompareErr{keys: []string{key}, err: err}
}
