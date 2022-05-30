package daos

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewTtDao,
)
