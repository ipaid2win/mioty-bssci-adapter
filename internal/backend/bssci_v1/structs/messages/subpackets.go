package messages

//go:generate msgp

// Subpackets
//
// reception info for every subpacket
type Subpackets struct {
	// Subpacket signal to noise ratio in dB
	SNR []int32 `msg:"snr"`
	// Subpacket signal strength in dBm
	RSSI []int32 `msg:"rssi"`
	// Subpacket frequencies in Hz
	Frequency []int32 `msg:"frequency"`
	// Subpacket phases in degree +-180, optional
	Phase []int32 `msg:"phase,omitempty"`
}
