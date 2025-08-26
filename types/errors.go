package types

import (
	"fmt"
	"strings"
)

type ResponseError struct {
	Status string `json:"status"`
	Err    struct {
		Code    string            `json:"code"`
		Message string            `json:"message"`
		Details map[string]string `json:"details"`
	} `json:"error"`
}

func (err ResponseError) Error() string {
	errStr := fmt.Sprintf("Error: status [%s] code [%s] message [%s] details [%s]", err.Status, err.Err.Code, err.Err.Message, mapToString(err.Err.Details))
	return errStr
}

func mapToString(m map[string]string) string {
	str := ""
	for k, v := range m {
		str += k + ":" + v + ","
	}
	return strings.Trim(str, ",")
}
