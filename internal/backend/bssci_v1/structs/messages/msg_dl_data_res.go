package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp
//msgp:replace common.EUI64 with:[8]byte

// Downlink data result
//
// The DL data result operation is initiated by the Base Station after queued DL data has
// either been sent or discarded.
//
// Basestation -> Service Center
type DlDataRes struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui"`
	// Assigned queue ID for reference, 64 bit
	QueId int64 `msg:"queId"`
	// sent, expired, invalid
	Result string `msg:"result"`
	// Unix UTC time of transmission, center of first subpacket, 64 bit, ns resolution, only if result is sent
	TxTime *uint64 `msg:"txTime"`
	// End Point packet counter, only if result is “sent”
	PacketCnt *uint32 `msg:"packetCnt"`
}

func NewDlDataRes(
	opId int64,
	epEui common.EUI64,
	queId int64,
	result string,
	txTime *uint64,
	packetCnt *uint32,
) DlDataRes {
	return DlDataRes{
		Command:   structs.MsgDlDataRes,
		OpId:      opId,
		EpEui:     epEui,
		QueId:     queId,
		Result:    result,
		TxTime:    txTime,
		PacketCnt: packetCnt,
	}
}

func (m *DlDataRes) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataRes) GetCommand() structs.Command {
	return structs.MsgDlDataRes
}

// Downlink data result response
//
// Service Center -> Basestation
type DlDataResRsp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
}

func NewDlDataResRsp(opId int64) DlDataResRsp {
	return DlDataResRsp{
		Command: structs.MsgDlDataResRsp,
		OpId:    opId,
	}
}

func (m *DlDataResRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataResRsp) GetCommand() structs.Command {
	return structs.MsgPing
}

// Downlink data result complete
//
// Basestation -> Service Center
type DlDataResCmp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
}

func NewDlDataResCmp(opId int64) DlDataResCmp {
	return DlDataResCmp{OpId: opId, Command: structs.MsgDlDataResCmp}
}

func (m *DlDataResCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataResCmp) GetCommand() structs.Command {
	return structs.MsgDlDataResCmp
}
