package v1

import (
	"github.com/google/wire"
)

type V1 struct {
	Tt *TtController
}

var ProviderSet = wire.NewSet(
	NewTtController,
)
