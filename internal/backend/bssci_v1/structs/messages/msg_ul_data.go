package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp
//msgp:replace common.EUI64 with:[8]byte

// Uplink data
//
// The UL data operation is initiated by the Base Station after receiving uplink data from
// an End Point. Telegrams carrying control data exclusively are considered as empty data.
//
// Basestation -> Service Center
type UlData struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui"`
	// Unix UTC time of reception, center of last subpacket, 64 bit, ns resolution
	RxTime uint64 `msg:"rxTime"`
	// Duration of the reception, center of first subpacket to center of last subpacket in ns, optional
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
	// Mioty mode and variant used for reception, i.e. ulp, ulp-rep, ulp-ll, optional
	Mode *string `msg:"mode,omitempty"`
	// Subpackets object with reception info for every subpacket, optional
	Subpackets *Subpackets `msg:"subpackets,omitempty"`
	// End Point user data, might be empty
	UserData []byte `msg:"userData"`
	// User data format identifier, 8 bit, optional, default 0
	Format *byte `msg:"format,omitempty"`
	// True if End Point downlink window is opened
	DlOpen bool `msg:"dlOpen"`
	// True if End Point expects a response in the DL window, requires dlOpen
	ResponseExp bool `msg:"responseExp"`
	// True if End Point acknowledges the reception of a DL transmission in the last DL window (packetCnt - 1)
	DlAck bool `msg:"dlAck"`
}

func NewUlData(
	opId int64,
	epEui common.EUI64,
	rxTime uint64,
	rxDuration *uint64,
	packetCnt uint32,
	snr float64,
	rssi float64,
	eqSnr *float64,
	profile *string,
	mode *string,
	subpackets *Subpackets,
	userData []byte,
	format *byte,
	dlOpen bool,
	responseExp bool,
	DlAck bool,

) UlData {
	return UlData{
		Command:     structs.MsgUlData,
		OpId:        opId,
		EpEui:       epEui,
		RxTime:      rxTime,
		RxDuration:  rxDuration,
		PacketCnt:   packetCnt,
		SNR:         snr,
		RSSI:        rssi,
		EqSnr:       eqSnr,
		Profile:     profile,
		Mode:        mode,
		Subpackets:  subpackets,
		UserData:    userData,
		Format:      format,
		DlOpen:      dlOpen,
		ResponseExp: responseExp,
		DlAck:       DlAck,
	}
}

func (m *UlData) GetOpId() int64 {
	return m.OpId
}

func (m *UlData) GetCommand() structs.Command {
	return structs.MsgUlData
}

// implements UplinkMessage.GetEndpointEui()
func (m *UlData) GetEndpointEui() common.EUI64 {
	return m.EpEui
}

// implements UplinkMessage.GetTsUnbRxInformation()
func (m *UlData) GetTsUnbRxInformation() UplinkMetadata {
	return UplinkMetadata{
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

// Uplink data response
//
// Service Center -> Basestation
type UlDataRsp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
}

func NewUlDataRsp(opId int64) UlDataRsp {
	return UlDataRsp{
		Command: structs.MsgUlDataRsp,
		OpId:    opId,
	}
}

func (m *UlDataRsp) GetOpId() int64 {
	return m.OpId
}

func (m *UlDataRsp) GetCommand() structs.Command {
	return structs.MsgPing
}

// UlDataach complete
//
// Basestation -> Service Center
type UlDataCmp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
}

func NewUlDataCmp(opId int64) UlDataCmp {
	return UlDataCmp{OpId: opId, Command: structs.MsgUlDataCmp}
}

func (m *UlDataCmp) GetOpId() int64 {
	return m.OpId
}

func (m *UlDataCmp) GetCommand() structs.Command {
	return structs.MsgUlDataCmp
}
