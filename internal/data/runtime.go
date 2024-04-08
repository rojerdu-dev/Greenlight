package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeForamt = errors.New("invalid runtime format")

type Runtime int32

func (r *Runtime) MarshalJSON() ([]byte, error) {
	// single-quoted string
	jsonValue := fmt.Sprintf("%d mins", r)

	// convert to double-quoted string for JSON
	quotedJSONValue := strconv.Quote(jsonValue)

	// convert to byte slice and return
	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeForamt
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeForamt
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeForamt
	}

	*r = Runtime(i)

	return nil
}
