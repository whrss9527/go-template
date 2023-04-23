package mysql

import "github.com/google/wire"

// ProviderSet is mysql repo providers.
var ProviderSet = wire.NewSet(NewTemplateRepo)
