package graph

import "github.com/claudioluciano/steamir/api/internal/pkg/steamclient"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*steamclient.SteamClient
}
