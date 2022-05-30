package service

import (
	"github.com/google/wire"
)

type Service struct {
	Tt *TtService
}

var ProviderSet = wire.NewSet(
	NewTtService,
)
