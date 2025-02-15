package bssci_v1

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

type connection struct {
	sync.RWMutex
	conn       net.Conn
	// stats      *stats.Collector
	lastActive time.Time
	version    string
	opId       int64
	// Base Station session UUID, used to resume session
	SnBsUuid uuid.UUID
	// Service Center session UUID, used to resume session
	SnScUuid uuid.UUID
}

// Send the message to this connection
func (conn *connection) Write(msg messages.Message, timeout time.Duration) (err error) {
	conn.Lock()
	defer conn.Unlock()

	bb, err := msg.MarshalMsg(nil)
	if err != nil {
		return errors.Wrap(err, "marshal msgp error")
	}

	conn.conn.SetWriteDeadline(time.Now().Add(timeout))
	_, err = conn.conn.Write(bb)
	if err != nil {
		conn.conn.Close()
		return
	}
	conn.lastActive = time.Now()

	return
}

// Read a message from this connection
func (conn *connection) Read(timeout time.Duration) (cmd structs.CommandHeader, raw msgp.Raw, err error) {
	conn.Lock()
	defer conn.Unlock()

	conn.conn.SetReadDeadline(time.Now().Add(timeout))
	cmd, raw, err = ReadBssciMessage(conn.conn)
	if err != nil {
		conn.conn.Close()
		return
	}
	conn.lastActive = time.Now()

	return
}

// Should be called when a message chain is initialized by the server.
//
// returns the current opId before decrement by 1
func (conn *connection) DecrementOpId() (opId int64) {
	conn.Lock()
	defer conn.Unlock()

	opId = conn.opId
	conn.opId = conn.opId - 1
	return
}

func (conn *connection) GetLastActive() time.Time {
	conn.RLock()
	defer conn.RUnlock()

	return conn.lastActive
}

// Check if this connection is resumed after a Con message
//
// returns true and the current snScUuid if the conenction is resumable, else false and a new snScUuid
func (conn *connection) ResumeConnection(snBsUuid uuid.UUID, snScOpId *int64) (resume bool, snScUuid uuid.UUID) {
	conn.Lock()
	defer conn.Unlock()

	conn.opId = -1

	if conn.SnBsUuid == snBsUuid {
		if snScOpId != nil {
			conn.opId = *snScOpId
		}
		return true, conn.SnScUuid
	}
	snScUuid = uuid.New()
	conn.SnScUuid = snScUuid
	
	return false, snScUuid

}
