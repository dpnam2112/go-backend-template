package dto

// APIResponse is a generic response structure
type APIResponse[T any] struct {
	Status int `json:"status"`         // HTTP Status Code
	Data   *T  `json:"data,omitempty"` // Data payload (generic)
}
