package data

import (
	"fmt"
	"strconv"
)

type runtime int32

func (r runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	jsonValue = strconv.Quote(jsonValue)
	return []byte(jsonValue), nil
}
