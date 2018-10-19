package googlepasses

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Header  http.Header `json:"-"`
	Errors  []ErrorItem `json:"errors"`
}

type ErrorItem struct {
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	if len(e.Errors) == 0 && e.Message == "" {
		return fmt.Sprintf("HTTP response code %d", e.Code)
	}

	errorString := fmt.Sprintf("Error %d: %s\n", e.Code, e.Message)
	for _, v := range e.Errors {
		errorString += "Reason: " + v.Reason + ", Message: " + v.Message + "\n"
	}

	return errorString
}
