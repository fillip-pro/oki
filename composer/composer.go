package composer

import "gitlab.com/fillip/oki/providers"

var (
	clients = providers.Providers()
)

// Composer provides the interface for composing new
// infrastructure pieces.
type Composer interface {
	Compose()
}
