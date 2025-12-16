package models

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}
