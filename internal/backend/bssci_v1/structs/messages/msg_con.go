package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"

	"github.com/google/uuid"
	// "mioty-bssci-adapter/internal/common"
	// "mioty-bssci-adapter/internal/common"
	// "github.com/google/uuid"
)

//go:generate msgp

// Connect
//
// The connect operation is initiated by the Base Station immediately after establishing the
// network connection with the Service Center. No other operations may be started by
// either the Base Station or the Service Center until the connect operation is completed.
// The initial connect operation must use ID 0. This still applies if a previous session shall
// be resumed.
//
// Basestation -> Service Center
type Con struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
	// Requested protocol version, major.minor.patch
	Version string `msg:"version"`
	// Base Station EUI64
	BsEui int64 `msg:"bsEui"`
	// Vendor of the Base Station, optional
	Vendor *string `msg:"vendor,omitempty"`
	// Model of the Base Station, optional
	Model *string `msg:"model,omitempty"`
	// Name of the Base Station, optional
	Name *string `msg:"name,omitempty"`
	// Software version, optional
	SwVersion *string `msg:"swVersion,omitempty"`
	// Additional Base Station information object, might contain arbitrary key-value-pairs, optional
	Info map[string]interface{} `msg:"info,omitempty"`
	// True if Base Station is bidirectional
	Bidi bool `msg:"bidi"`
	// Geographic location [Latitude, Longitude, Altitude], optional
	GeoLocation *GeoLocation `msg:"geoLocation,omitempty"`
	// Base Station session UUID, must match with previous connect to resume session
	SnBsUuid structs.SessionUuid `msg:"snBsUuid"`
	// Minimum required known Base Station operation ID to resume previous session, optional
	SnBsOpId *int64 `msg:"snBsOpId,omitempty"`
	// Maximum known Service Center operation ID to resume previous session, optional
	SnScOpId *int64 `msg:"snScOpId,omitempty"`
}

func (m *Con) GetOpId() int64 {
	return m.OpId
}

func (m *Con) GetCommand() structs.Command {
	return structs.MsgCon
}

func (m *Con) GetEui() common.EUI64 {
	return int64ToEui(m.BsEui)
}

// Connect response
//
// Basestation -> Service Center
type ConRsp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
	// Requested protocol version, major.minor.patch
	Version string `msg:"version"`
	// Service Center EUI64
	ScEui int64 `msg:"scEui"`
	// Vendor of the Service Center, optional
	Vendor *string `msg:"vendor,omitempty"`
	// Model of the Service Center, optional
	Model *string `msg:"model,omitempty"`
	// Name of the Service Center, optional
	Name *string `msg:"name,omitempty"`
	// Software version, optional
	SwVersion *string `msg:"swVersion,omitempty"`
	// Additional Service Center information object, might contain arbitrary key-value-pairs, optional
	Info map[string]interface{} `msg:"info,omitempty"`
	// True if a previous session is resumed
	SnResume bool `msg:"snResume"`
	// Service Center session UUID, must match with previous connect to resume sessionF
	SnScUuid structs.SessionUuid `msg:"snScUuid"`
}

func NewConRsp(opId int64, version string, snScUuid uuid.UUID) ConRsp {
	session := structs.NewSessionUuid(snScUuid)
	vendor := "ChirpStackBssci"
	name := "ChirpStackBssci"
	model := "ChirpStackBssci"
	swVersion := "1.0"
	return ConRsp{
		Command:   structs.MsgConRsp,
		OpId:      opId,
		ScEui:     1,
		Version:   version,
		SnScUuid:  session,
		SnResume:  false,
		Vendor:    &vendor,
		Name:      &name,
		Model:     &model,
		SwVersion: &swVersion,
	}
}

func (m *ConRsp) ResumeConnection(snScUuid uuid.UUID) {
	session := structs.NewSessionUuid(snScUuid)
	m.SnScUuid = session
	m.SnResume = true
}

func (m *ConRsp) GetOpId() int64 {
	return m.OpId
}

func (m *ConRsp) GetCommand() structs.Command {
	return structs.MsgConRsp
}

func (m *ConRsp) GetEui() common.EUI64 {
	return int64ToEui(m.ScEui)
}

// Connect complete
//
// Basestation -> Service Center
type ConCmp struct {
	Command structs.Command `msg:"command"`
	// ID of the operation
	OpId int64 `msg:"opId"`
}

func NewConCmp(opId int64) ConCmp {
	return ConCmp{OpId: opId, Command: structs.MsgConCmp}
}

func (m *ConCmp) GetOpId() int64 {
	return m.OpId
}

func (m *ConCmp) GetCommand() structs.Command {
	return structs.MsgConCmp
}

// helper function to parse the int64 formatted eui64
func int64ToEui(in int64) common.EUI64 {

	b := [8]byte{
		byte(0xff & in),
		byte(0xff & (in >> 8)),
		byte(0xff & (in >> 16)),
		byte(0xff & (in >> 24)),
		byte(0xff & (in >> 32)),
		byte(0xff & (in >> 40)),
		byte(0xff & (in >> 48)),
		byte(0xff & (in >> 56)),
	}

	return b

}
