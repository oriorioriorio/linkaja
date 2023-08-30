package model

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type ServiceError struct {
	Code    int
	Message string
}

func (s ServiceError) Error() string {
	return s.Message
}
