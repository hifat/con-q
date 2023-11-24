package errorDomain

type Error struct {
	Status    int    `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	Code      string `json:"code,omitempty"`
	Attribute any    `json:"attribute,omitempty"`
}

type Response struct {
	Error Error `json:"error"`
}

func (e Error) Error() string {
	return e.Message
}
