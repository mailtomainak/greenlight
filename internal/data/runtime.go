package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type runtime int32

func (r runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	jsonValue = strconv.Quote(jsonValue)
	return []byte(jsonValue), nil
}

func (r *runtime) UnmarshalJSON(json []byte) error {

	jsonString, err := strconv.Unquote(string(json))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	splitJson := strings.Split(jsonString, " ")
	if len(splitJson) != 2 || splitJson[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(splitJson[0], 10, 32)

	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = runtime(i)
	return nil
}
