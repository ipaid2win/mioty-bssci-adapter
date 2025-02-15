package events

import "mioty-bssci-adapter/internal/common"

// Subscribe event
type Subscribe struct {
	// Gateway ID.
	GatewayEui common.EUI64

	// Subscribe (true) or unsubscribe (false) the gateway.
	Subscribe bool
}
