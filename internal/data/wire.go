package data

import (
	"github.com/google/wire"
	"tt.com/tt/internal/data/db"
)

var ProviderSet = wire.NewSet(
	db.NewRedis,
	db.NewDb,
	NewDataSource,
)
