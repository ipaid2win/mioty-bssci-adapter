package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"

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


type UplinkMessage interface {
	Message
	GetEndpointEui() common.EUI64
	GetUplinkMetadata() UplinkMetadata
	IntoProto(bsEui common.EUI64) (*msg.EndnodeUplink, error)
}

type UplinkMetadata struct {
	RxTime     uint64      `json:"rxTime"`
	RxDuration *uint64     `json:"rxDuration,omitempty"`
	PacketCnt  uint32      `json:"packetCnt"`
	Profile    *string     `json:"profile,omitempty"`
	SNR        float64     `json:"snr"`
	RSSI       float64     `json:"rssi"`
	EqSnr      *float64    `json:"eqSnr,omitempty"`
	Subpackets *Subpackets `json:"subpackets,omitempty"`
}

type PropagateMessage interface {
	Message
	GetEndpointEui() common.EUI64
	IntoProto() cmd.PropagateEndnode
}

