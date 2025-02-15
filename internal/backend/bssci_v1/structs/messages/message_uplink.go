package messages

import (
	"mioty-bssci-adapter/internal/common"

)

type UplinkMessage interface {
	Message
	GetEndpointEui() common.EUI64
	GetTsUnbInformation() TsUnbInformation
}

type TsUnbInformation struct {
	RxTime     uint64      `json:"rxTime"`
	RxDuration *uint64     `json:"rxDuration,omitempty"`
	PacketCnt  uint32      `json:"packetCnt"`
	Profile    *string     `json:"profile,omitempty"`
	EqSnr      *float64    `json:"eqSnr,omitempty"`
	SNR        float64     `json:"snr"`
	RSSI       float64     `json:"rssi"`
	Subpackets *Subpackets `json:"subpackets,omitempty"`
}

