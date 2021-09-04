package common

const (
	status_error   = "error"
	status_fail    = "fail"
	status_success = "success"
)

type body struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func ResponseSuccess(data interface{}) body {
	return body{
		Status: status_success,
		Data:   data,
	}
}

func ResponseFail(data interface{}) body {
	return body{
		Status: status_fail,
		Data:   data,
	}
}

func ResponseError(message string, code int, data interface{}) body {
	return body{
		Status:  status_error,
		Message: message,
		Code:    code,
		Data:    data,
	}
}