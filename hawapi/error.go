package hawapi

import "fmt"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Method  string `json:"method"`
	Cause   string `json:"cause"`
	Url     string `json:"url"`
	Message string `json:"message,omitempty"`
}

func (e ErrorResponse) Error() string {
	msg := fmt.Sprintf("request error [%s %d] using %s method", e.Status, e.Code, e.Method)

	if len(e.Url) != 0 {
		msg = fmt.Sprintf("%s on '%s'", msg, e.Url)
	}

	if len(e.Message) != 0 {
		msg = fmt.Sprintf("%s: %s", msg, e.Message)
	}

	return msg
}
