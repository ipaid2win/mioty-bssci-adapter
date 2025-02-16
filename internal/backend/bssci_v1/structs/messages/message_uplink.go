package messages

import (
	"mioty-bssci-adapter/internal/common"
)

type UplinkMessage interface {
	Message
	GetEndpointEui() common.EUI64
	GetUplinkMetadata() UplinkMetadata
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
