package utils

// ApplicationErr to store error information
type ApplicationErr struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
