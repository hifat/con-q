package httpDomain

type ResSucces[T any] struct {
	Item    *T       `json:"item,omitempty"`
	Items   []T      `json:"items,omitempty"`
	Message string   `json:"message,omitempty"`
	Meta    *ResMeta `json:"meta,omitempty"`
}

type Pagination struct {
	Total   *int64 `json:"total,omitempty"`
	Page    *int   `json:"page,omitempty"`
	PerPage *int   `json:"perPage,omitempty"`
}

type ResMeta struct {
	Pagination *Pagination `json:"pagination"`
}
