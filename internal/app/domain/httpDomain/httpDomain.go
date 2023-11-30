package httpDomain

type ResSucces struct {
	Item    any     `json:"item,omitempty"`
	Items   any     `json:"items,omitempty"`
	Message string  `json:"message,omitempty"`
	Meta    ResMeta `json:"meta,omitempty"`
}

type ResMeta struct {
	Total   uint `json:"total,omitempty"`
	Page    uint `json:"page,omitempty"`
	PerPage uint `json:"perPage,omitempty"`
}
