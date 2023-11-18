package httpDomain

type SuccesResponse struct {
	Item    any    `json:"item,omitempty"`
	Items   []any  `json:"items,omitempty"`
	Total   int    `json:"total,omitempty"`
	Message string `json:"message,omitempty"`
}
