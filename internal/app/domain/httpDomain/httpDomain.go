package httpDomain

type ResSucces struct {
	Item    any      `json:"item,omitempty"`
	Items   any      `json:"items,omitempty"`
	Message string   `json:"message,omitempty"`
	Meta    *ResMeta `json:"meta,omitempty"`
}

type ResMeta struct {
	Total   *int `json:"total,omitempty"`
	Page    *int `json:"page,omitempty"`
	PerPage *int `json:"perPage,omitempty"`
}
