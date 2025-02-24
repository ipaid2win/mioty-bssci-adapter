package structs

//go:generate msgp
//msgp:shim uuid.UUID as:SessionUuid using:(SessionUuid).ToUuid/NewSessionUuid

import (
	"github.com/google/uuid"
)

// BSSCI encodes as 16 signed ints
type SessionUuid [16]int8

func NewSessionUuid(uuid uuid.UUID) (s SessionUuid) {
	for i, v := range uuid {
		s[i] = int8(v)
	}
	return
}

func (s SessionUuid) ToUuid() (u uuid.UUID) {
	for i, v := range s {
		u[i] = byte(v)
	}
	return
}

func (s SessionUuid) String() string {
	return s.ToUuid().String()
}

// MarshalText implements encoding.TextMarshaler.
func (s SessionUuid) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *SessionUuid) UnmarshalText(text []byte) error {
	u, err := uuid.ParseBytes(text)
	if err != nil {
		return err
	}

	for i, v := range u {
		s[i] = int8(v)
	}

	return nil
}

