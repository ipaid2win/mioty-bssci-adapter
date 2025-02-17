package bssci_v1

import (
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"sync"

	"github.com/pkg/errors"
)

var (
	errGatewayDoesNotExist = errors.New("gateway does not exist")
)

type basestations struct {
	sync.RWMutex
	basestations          map[common.EUI64]*connection
	subscribeEventHandler func(events.Subscribe)
}

func (g *basestations) get(eui common.EUI64) (*connection, error) {
	g.RLock()
	defer g.RUnlock()

	gw, ok := g.basestations[eui]
	if !ok {
		return gw, errGatewayDoesNotExist
	}
	return gw, nil
}

func (g *basestations) set(eui common.EUI64, c *connection) error {
	g.Lock()
	defer g.Unlock()

	g.basestations[eui] = c

	if g.subscribeEventHandler != nil {
		g.subscribeEventHandler(events.Subscribe{Subscribe: true, GatewayEui: eui})
	}

	return nil
}

func (g *basestations) remove(eui common.EUI64) error {
	g.Lock()
	defer g.Unlock()

	if g.subscribeEventHandler != nil {
		g.subscribeEventHandler(events.Subscribe{Subscribe: false, GatewayEui: eui})
	}

	delete(g.basestations, eui)
	return nil
}
