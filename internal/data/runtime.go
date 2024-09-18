package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntime = fmt.Errorf("invalid runtime format ")

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJson := strconv.Quote(jsonValue)

	return []byte(quotedJson), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJson, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntime
	}
	parts := strings.Split(unquotedJson, " ") // split the int value
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntime
	}
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntime
	}
	*r = Runtime(i)
	return nil
}
