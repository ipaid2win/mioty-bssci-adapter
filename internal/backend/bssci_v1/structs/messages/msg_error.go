package messages

import "mioty-bssci-adapter/internal/backend/bssci_v1/structs"

//go:generate msgp

// Error
//
// An error message might be send in any operation in case of an error condition. The
// error message terminates the regular operation sequence of initiation, response and
// completion after either the initiation or the response, depending on where the error
// condition occurs. In both cases the operation then follows the sequence of error and
// error acknowledgement instead, with the error acknowledgement completing the
// operation.
//
// Service Center <-> Basestation
type BssciError struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// Error code, using POSIX error numbers,
	Code uint32 `msg:"code" json:"code"`
	// Error message
	Message string `msg:"message" json:"message"`
}

func NewBssciError(opId int64, code uint32, message string) BssciError {
	return BssciError{OpId: opId, Command: structs.MsgError, Code: code, Message: message}
}

func (m *BssciError) GetOpId() int64 {
	return m.OpId
}

func (m *BssciError) GetCommand() structs.Command {
	return structs.MsgError
}

// Error Ack
//
// Basestation <-> Service Center
type BssciErrorAck struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewBssciErrorAck(opId int64) BssciErrorAck {
	return BssciErrorAck{
		Command: structs.MsgErrorAck,
		OpId:    opId,
	}
}

func (m *BssciErrorAck) GetOpId() int64 {
	return m.OpId
}

func (m *BssciErrorAck) GetCommand() structs.Command {
	return structs.MsgErrorAck
}
