package bssci_v1

import (
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	errGatewayDoesNotExist = errors.New("gateway does not exist")
)


type basestations struct {
	sync.RWMutex
	basestations       map[common.EUI64]*connection
	subscribeEventFunc func(events.Subscribe)
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

	if g.subscribeEventFunc != nil {
		g.subscribeEventFunc(events.Subscribe{Subscribe: true, GatewayEui: eui})
	}

	return nil
}

func (g *basestations) getLastActive(eui common.EUI64) (time.Time, error) {
	g.RLock()
	defer g.RUnlock()

	gw, ok := g.basestations[eui]
	if !ok {
		return time.Time{}, errGatewayDoesNotExist
	}

	return gw.lastActive, nil
}

func (g *basestations) setLastTimesync(eui common.EUI64, ts time.Time) error {
	g.Lock()
	defer g.Unlock()

	gw, ok := g.basestations[eui]
	if !ok {
		return errGatewayDoesNotExist
	}

	gw.lastActive = ts
	g.basestations[eui] = gw

	return nil
}



func (g *basestations) remove(eui common.EUI64) error {
	g.Lock()
	defer g.Unlock()

	if g.subscribeEventFunc != nil {
		g.subscribeEventFunc(events.Subscribe{Subscribe: false, GatewayEui: eui})
	}

	delete(g.basestations, eui)
	return nil
}
