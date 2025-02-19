package messages

import (
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"

	"github.com/tinylib/msgp/msgp"
)

// Each message must implement this
type Message interface {
	// get the opId
	GetOpId() int64
	// get the name of this message type
	GetCommand() structs.Command
	// message pack interfaces
	msgp.Encodable
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Decodable
}

type EndnodeUplinkMessage interface {
	Message
	GetEndpointEui() common.EUI64
	GetUplinkMetadata() UplinkMetadata
	IntoProto(bsEui common.EUI64) (*msg.EndnodeUplink, error)
}

type PropagateMessage interface {
	Message
	GetEndpointEui() common.EUI64
	IntoProto() cmd.PropagateEndnode
}

type BasestationStatusMessage interface {
	Message
	IntoProto(bsEui common.EUI64) (*msg.BasestationStatus, error)
}