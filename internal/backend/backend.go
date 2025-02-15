package backend

import (
	"fmt"
	"mioty-bssci-adapter/internal/backend/bssci_v1"
	"mioty-bssci-adapter/internal/config"

	"github.com/pkg/errors"
)



var backend Backend

// Setup configures the backend.
func Setup(conf config.Config) error {
	var err error

	switch conf.Backend.Type {
	case "bssci_v1":
		backend, err = bssci_v1.NewBackend(conf)
	default:
		return fmt.Errorf("unknown backend type: %s", conf.Backend.Type)
	}

	if err != nil {
		return errors.Wrap(err, "new backend error")
	}

	return nil
}

// GetBackend returns the backend.
func GetBackend() Backend {
	return backend
}

// Backend defines the interface that a backend must implement
type Backend interface {
	// Stop closes the backend.
	Stop() error

	// Start starts the backend.
	Start() error


}
