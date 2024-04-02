package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	// single-quoted string
	jsonValue := fmt.Sprintf("%d mins", r)

	// convert to double-quoted string for JSON
	quotedJSONValue := strconv.Quote(jsonValue)

	// convert to byte slice and return
	return []byte(quotedJSONValue), nil
}
