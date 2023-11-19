package httpDomain

type SuccesResponse[T any] struct {
	Item    T      `json:"item,omitempty"`
	Items   []T    `json:"items,omitempty"`
	Total   int    `json:"total,omitempty"`
	Message string `json:"message,omitempty"`
}
