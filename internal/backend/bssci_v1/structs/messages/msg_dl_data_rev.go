package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp
//msgp:shim common.EUI64 as:int64 using:common.Eui64toInt/common.Eui64FromInt

// Downlink data revoke
//
// The DL data revoke operation is initiated by the Service Center to revoke previously
// scheduled downlink data at the Base Station.
//
// Service Center -> Basestation
type DlDataRev struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// Assigned queue ID for reference, 64 bit
	QueId uint64 `msg:"queId" json:"queId"`
}

func NewDlDataRev(
	opId int64,
	epEui common.EUI64,
	queId uint64,
) DlDataRev {
	return DlDataRev{
		Command: structs.MsgDlDataRev,
		OpId:    opId,
		EpEui:   epEui,
		QueId:   queId,
	}
}

func (m *DlDataRev) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataRev) GetCommand() structs.Command {
	return structs.MsgDlDataRev
}

// Downlink data revoke response
//
// Basestation -> Service Center
type DlDataRevRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlDataRevRsp(opId int64) DlDataRevRsp {
	return DlDataRevRsp{
		Command: structs.MsgDlDataRevRsp,
		OpId:    opId,
	}
}

func (m *DlDataRevRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataRevRsp) GetCommand() structs.Command {
	return structs.MsgPing
}

// Downlink data revoke complete
//
// Service Center -> Basestation
type DlDataRevCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlDataRevCmp(opId int64) DlDataRevCmp {
	return DlDataRevCmp{OpId: opId, Command: structs.MsgDlDataRevCmp}
}

func (m *DlDataRevCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataRevCmp) GetCommand() structs.Command {
	return structs.MsgDlDataRevCmp
}
