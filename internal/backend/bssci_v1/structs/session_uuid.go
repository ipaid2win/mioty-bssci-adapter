package structs

//go:generate msgp
//msgp:shim uuid.UUID as:SessionUuid using:(SessionUuid).ToUuid/NewSessionUuid

import (
	"github.com/google/uuid"
)

// BSSCI encodes as 16 signed ints
type SessionUuid [16]int8

func NewSessionUuid(uuid uuid.UUID) (session SessionUuid) {
	for i, s := range uuid {
		session[i] = int8(s)
	}
	return
}

func (session SessionUuid) ToUuid() (u uuid.UUID) {
	for i, s := range session {
		u[i] = byte(s)
	}
	return
}

func (session SessionUuid) String() string {
	return session.ToUuid().String()
}
