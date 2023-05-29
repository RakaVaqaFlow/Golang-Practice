package graph

import (
	"workshop-8-3/bff/internal/downstreams/api"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoAPI *api.Client
}
