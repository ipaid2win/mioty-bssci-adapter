package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp
//msgp:replace common.EUI64 with:[8]byte

// Detach
//
// The detach operation is initiated by the Base Station after receiving an over the air
// detachment request from an End Point.
//
// Basestation -> Service Center
type Det struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui"`
	// Unix UTC time of reception, center of last subpacket, 64 bit, ns resolution
	RxTime uint64 `msg:"rxTime"`
	// Duration or the reception, center of first subpacket to center of last subpacket in ns, optional
	RxDuration *uint64 `msg:"rxDuration,omitempty"`
	// End Point packet counter
	PacketCnt uint32 `msg:"packetCnt"`
	// Reception signal to noise ratio in dB
	SNR float64 `msg:"snr"`
	// Reception signal strength in dBm
	RSSI float64 `msg:"rssi"`
	// AWGN equivalent reception SNR in dB, optional
	EqSnr *float64 `msg:"eqSnr,omitempty"`
	// Name of the Mioty profile used for reception, i.e. eu1, optional
	Profile *string `msg:"profile,omitempty"`
	// Subpackets object with reception info for every subpacket, optional
	Subpackets *Subpackets `msg:"subpackets,omitempty"`
	// End Point signature
	Sign [4]byte `msg:"sign"`
}

func NewDet(
	opId int64,
	epEui common.EUI64,
	rxTime uint64,
	rxDuration *uint64,
	packetCnt uint32,
	snr float64,
	rssi float64,
	eqSnr *float64,
	profile *string,
	subpackets *Subpackets,
	sign [4]byte,
) Det {
	return Det{
		Command:    structs.MsgDet,
		OpId:       opId,
		EpEui:      epEui,
		RxTime:     rxTime,
		RxDuration: rxDuration,
		PacketCnt:  packetCnt,
		RSSI:       rssi,
		SNR:        snr,
		Profile:    profile,
		Subpackets: subpackets,
		Sign:       sign,
	}
}

func (m *Det) GetOpId() int64 {
	return m.OpId
}

func (m *Det) GetCommand() structs.Command {
	return structs.MsgDet
}

// respond to Det with DetRsp
func (m *Det) Respond() Message {
	msg := NewDetRsp(m.OpId, m.Sign)
	return &msg
}

// implements UplinkMessage.GetEndpointEui()
func (m *Det) GetEndpointEui() common.EUI64 {
	return m.EpEui
}

// implements UplinkMessage.GetTsUnbRxInformation()
func (m *Det) GetTsUnbRxInformation() TsUnbInformation {
	return TsUnbInformation{
		RxTime:     m.RxTime,
		RxDuration: m.RxDuration,
		PacketCnt:  m.PacketCnt,
		Profile:    m.Profile,
		SNR:        m.SNR,
		RSSI:       m.RSSI,
		EqSnr:      m.EqSnr,
		Subpackets: m.Subpackets,
	}
}



// Detach response
//
// Service Center -> Basestation
type DetRsp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
	// End Point signature
	Sign [4]byte `msg:"sign"`
}

func NewDetRsp(opId int64, sign [4]byte) DetRsp {
	return DetRsp{
		Command: structs.MsgDetRsp,
		OpId:    opId,
		Sign:    sign,
	}
}

func (m *DetRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DetRsp) GetCommand() structs.Command {
	return structs.MsgPing
}

// Detach complete
//
// Basestation -> Service Center
type DetCmp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
}

func NewDetCmp(opId int64) DetCmp {
	return DetCmp{OpId: opId, Command: structs.MsgDetCmp}
}

func (m *DetCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DetCmp) GetCommand() structs.Command {
	return structs.MsgDetCmp
}
