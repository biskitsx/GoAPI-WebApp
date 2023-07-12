package utils

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
}

func CreateError(statusCode int, err interface{}) ErrorResponse {

	return ErrorResponse{
		StatusCode: statusCode,
		Error:      err,
	}

}
