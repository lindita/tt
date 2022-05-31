package h5

import (
	"github.com/google/wire"
)

type H5 struct {
	Tt *ttController
}

var ProviderSet = wire.NewSet(
	NewTtController,
)
