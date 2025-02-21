package messages

import (
	"mioty-bssci-adapter/internal/api/msg"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func (m *UplinkMetadata) IntoProto() *msg.EndnodeUplinkMetadata {
	var message msg.EndnodeUplinkMetadata
	if m != nil {
		rxTime := timestamppb.Timestamp{
			Seconds: int64(m.RxTime / 1000),
			Nanos:   int32(m.RxTime % 1000),
		}
		message = msg.EndnodeUplinkMetadata{
			RxTime:    &rxTime,
			PacketCnt: m.PacketCnt,
			Profile:   m.Profile,
			Rssi:      m.RSSI,
			Snr:       m.SNR,
			EqSnr:     m.EqSnr,
		}
		if m.Subpackets != nil {
			message.SubpacketInfo = m.Subpackets.IntoProto()
		}
		if m.RxDuration != nil {
			rxDuration := durationpb.Duration{
				Seconds: int64(*m.RxDuration / 1000),
				Nanos:   int32(*m.RxDuration % 1000),
			}
			message.RxDuration = &rxDuration
		}
	}
	return &message
}
