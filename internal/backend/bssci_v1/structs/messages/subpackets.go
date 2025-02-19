package messages

import "mioty-bssci-adapter/internal/api/msg"

//go:generate msgp

// Subpackets
//
// reception info for every subpacket
type Subpackets struct {
	// Subpacket signal to noise ratio in dB
	SNR []int32 `msg:"snr" json:"snr"`
	// Subpacket signal strength in dBm
	RSSI []int32 `msg:"rssi" json:"rssi"`
	// Subpacket frequencies in Hz
	Frequency []int32 `msg:"frequency" json:"frequency"`
	// Subpacket phases in degree +-180, optional
	Phase []int32 `msg:"phase,omitempty" json:"phase,omitempty"`
}

func (subpackets *Subpackets) IntoProto() []*msg.EndnodeUplinkSubpacket {

	pb := make([]*msg.EndnodeUplinkSubpacket, 0, len(subpackets.RSSI))

	if subpackets.Phase == nil {
		for i := 1; i < len(subpackets.RSSI); i++ {
			proto := msg.EndnodeUplinkSubpacket{
				Snr:       subpackets.SNR[i],
				Rssi:      subpackets.RSSI[i],
				Frequency: subpackets.Frequency[i],
			}
			pb = append(pb, &proto)
		}

	} else {
		for i := 1; i < len(subpackets.RSSI); i++ {
			proto := msg.EndnodeUplinkSubpacket{
				Snr:       subpackets.SNR[i],
				Rssi:      subpackets.RSSI[i],
				Frequency: subpackets.Frequency[i],
				Phase:     &subpackets.Phase[i],
			}
			pb = append(pb, &proto)
		}
	}

	return pb
}
