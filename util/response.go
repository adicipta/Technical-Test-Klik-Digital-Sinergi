package util

const (
	statusSuccess = "sucess"
	statusFail    = "fail"
	statusError   = "error"
)

type body struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSucess(message string, data interface{}) body {
	return body{
		Status:  statusSuccess,
		Message: message,
		Data:    data,
	}
}

func ResponseFail(message string, data interface{}) body {
	return body{
		Status:  statusFail,
		Message: message,
		Data:    data,
	}
}

func ResponseError(message string, data interface{}) body {
	return body{
		Status:  statusError,
		Message: message,
		Data:    data,
	}
}
