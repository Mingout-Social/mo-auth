package lib

type SuccessResponse struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type ErrorResponse struct {
	Error        bool   `json:"error"`
	ErrorMessage string `json:"error_message"`
}
