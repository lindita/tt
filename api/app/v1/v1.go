package v1

import (
	"github.com/google/wire"
)

type V1 struct {
	Tt *ttController
}

var ProviderSet = wire.NewSet(
	NewTtController,
)
