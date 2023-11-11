package di

import (
	"github.com/hifat/con-q/internal/app/handler"

	"github.com/google/wire"
)

var AdapterSet = wire.NewSet(NewAdapter)

type Adapter struct {
	Handler handler.Handler
}

func NewAdapter(h handler.Handler) Adapter {
	return Adapter{
		Handler: h,
	}
}
