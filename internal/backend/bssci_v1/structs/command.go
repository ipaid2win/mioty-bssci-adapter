package structs

//go:generate msgp

type CommandHeader struct {
	Command Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
}

func (m *CommandHeader) GetCommand() Command {
	return m.Command
}

func (m *CommandHeader) GetOpId() int64 {
	return m.OpId
}

// Command defines the message type.
type Command string

// Message types.
const (
	MsgAtt            Command = "att"
	MsgAttRsp         Command = "attRsp"
	MsgAttCmp         Command = "attCmp"
	MsgAttPrp         Command = "attPrp"
	MsgAttPrpRsp      Command = "attPrpRsp"
	MsgAttPrpCmp      Command = "attPrpCmp"
	MsgCon            Command = "con"
	MsgConRsp         Command = "conRsp"
	MsgConCmp         Command = "conCmp"
	MsgDet            Command = "det"
	MsgDetRsp         Command = "detRsp"
	MsgDetCmp         Command = "detCmp"
	MsgDetPrp         Command = "detPrp"
	MsgDetPrpRsp      Command = "detPrpRsp"
	MsgDetPrpCmp      Command = "detPrpCmp"
	MsgDlDataQue      Command = "dlDataQue"
	MsgDlDataQueRsp   Command = "dlDataQueRsp"
	MsgDlDataQueCmp   Command = "dlDataQueCmp"
	MsgDlDataRes      Command = "dlDataRes"
	MsgDlDataResRsp   Command = "dlDataResRsp"
	MsgDlDataResCmp   Command = "dlDataResCmp"
	MsgDlDataRev      Command = "dlDataRev"
	MsgDlDataRevRsp   Command = "dlDataRevRsp"
	MsgDlDataRevCmp   Command = "dlDataRevCmp"
	MsgDlRxStat       Command = "dlRxStat"
	MsgDlRxStatRsp    Command = "dlRxStatRsp"
	MsgDlRxStatCmp    Command = "dlRxStatCmp"
	MsgDlRxStatQry    Command = "dlRxStatQry"
	MsgDlRxStatQryRsp Command = "dlRxStatQryRsp"
	MsgDlRxStatQryCmp Command = "dlRxStatQryCmp"
	MsgPing           Command = "ping"
	MsgPingRsp        Command = "pingRsp"
	MsgPingCmp        Command = "pingCmp"
	MsgStatus         Command = "status"
	MsgStatusRsp      Command = "statusRsp"
	MsgStatusCmp      Command = "statusCmp"
	MsgUlData         Command = "ulData"
	MsgUlDataRsp      Command = "ulDataRsp"
	MsgUlDataCmp      Command = "ulDataCmp"
	MsgError          Command = "error"
	MsgErrorAck       Command = "errorAck"
)

// A message send by the server
const (
	ServerMsgAttPrp         Command = MsgAttPrp
	ServerMsgAttPrpCmp      Command = MsgAttPrpCmp
	ServerMsgDetPrp         Command = MsgDetPrp
	ServerMsgDetPrpCmp      Command = MsgDetPrpCmp
	ServerMsgDlDataQue      Command = MsgDlDataQue
	ServerMsgDlDataQueCmp   Command = MsgDlDataQueCmp
	ServerMsgDlDataRev      Command = MsgDlDataRev
	ServerMsgDlDataRevCmp   Command = MsgDlDataRevCmp
	ServerMsgDlRxStatQry    Command = MsgDlRxStatQry
	ServerMsgDlRxStatQryCmp Command = MsgDlRxStatQryCmp
	ServerMsgStatus         Command = MsgStatus
	ServerMsgStatusCmp      Command = MsgStatusCmp
	ServerMsgAttRsp         Command = MsgAttRsp
	ServerMsgDetRsp         Command = MsgDetRsp
	ServerMsgConRsp         Command = MsgConRsp
	ServerMsgDlDataResRsp   Command = MsgDlDataResRsp
	ServerMsgUlDataRsp      Command = MsgUlDataRsp
	ServerMsgDlRxStatRsp    Command = MsgDlRxStatRsp
	ServerMsgError          Command = MsgError
	ServerMsgErrorAck       Command = MsgErrorAck
	ServerMsgPing           Command = MsgPing
	ServerMsgPingRsp        Command = MsgPingRsp
	ServerMsgPingCmp        Command = MsgPingCmp
)

// A message send by the client
const (
	ClientMsgAtt            Command = MsgAtt
	ClientMsgAttCmp         Command = MsgAttCmp
	ClientMsgAttPrpRsp      Command = MsgAttPrpRsp
	ClientMsgCon            Command = MsgCon
	ClientMsgConCmp         Command = MsgConCmp
	ClientMsgDet            Command = MsgDet
	ClientMsgDetCmp         Command = MsgDetCmp
	ClientMsgDetPrpRsp      Command = MsgDetPrpRsp
	ClientMsgDlDataQueRsp   Command = MsgDlDataQueRsp
	ClientMsgDlDataRes      Command = MsgDlDataRes
	ClientMsgDlDataResCmp   Command = MsgDlDataResCmp
	ClientMsgDlDataRevRsp   Command = MsgDlDataRevRsp
	ClientMsgDlRxStat       Command = MsgDlRxStat
	ClientMsgDlRxStatCmp    Command = MsgDlRxStatCmp
	ClientMsgDlRxStatQryRsp Command = MsgDlRxStatQryRsp
	ClientMsgStatusRsp      Command = MsgStatusRsp
	ClientMsgUlData         Command = MsgUlData
	ClientMsgUlDataCmp      Command = MsgUlDataCmp
	ClientMsgError          Command = MsgError
	ClientMsgErrorAck       Command = MsgErrorAck
	ClientMsgPing           Command = MsgPing
	ClientMsgPingRsp        Command = MsgPingRsp
	ClientMsgPingCmp        Command = MsgPingCmp
)
