package messages

import "mioty-bssci-adapter/internal/backend/bssci_v1/structs"

//go:generate msgp

// Ping
//
// The ping operation can be initiated by either the Base Station or the Service Center to
// verify an established connection during idle times where no other operations are
// initiated.
//
// Service Center <-> Basestation
type Ping struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewPing(opId int64) Ping {
	return Ping{OpId: opId, Command: structs.MsgPing}
}

func (m *Ping) GetOpId() int64 {
	return m.OpId
}

func (m *Ping) GetCommand() structs.Command {
	return structs.MsgPing
}

// Ping response
//
// Basestation <-> Service Center
type PingRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewPingRsp(opId int64) PingRsp {
	return PingRsp{OpId: opId, Command: structs.MsgPingRsp}
}

func (m *PingRsp) GetOpId() int64 {
	return m.OpId
}

func (m *PingRsp) GetCommand() structs.Command {
	return structs.MsgPingRsp
}

// Ping complete
//
// Service Center <-> Basestation
type PingCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewPingCmp(opId int64) PingCmp {
	return PingCmp{OpId: opId, Command: structs.MsgPingCmp}
}

func (m *PingCmp) GetOpId() int64 {
	return m.OpId
}

func (m *PingCmp) GetCommand() structs.Command {
	return structs.MsgPingCmp
}
