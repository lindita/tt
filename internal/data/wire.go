package data

import (
	"github.com/google/wire"
	"test.com/tt/internal/data/db"
)

var ProviderSet = wire.NewSet(
	db.NewRedis,
	db.NewDb,
	NewDataSource,
)
