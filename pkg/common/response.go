package common

type Response[T any] struct {
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Status  uint16 `json:"status"`
	Success bool   `json:"success"`
}
