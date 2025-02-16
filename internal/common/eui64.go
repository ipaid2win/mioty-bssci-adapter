package common

import (
	"encoding/hex"
	"fmt"
	"strings"
)


// EUI64 data type
type EUI64 [8]byte

// helper function to parse a int64 formatted eui64
func EUI64FromInt(in int64) EUI64 {
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


// MarshalText implements encoding.TextMarshaler.
func (e EUI64) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *EUI64) UnmarshalText(text []byte) error {
	b, err := hex.DecodeString(strings.TrimPrefix(string(text), "0x"))
	if err != nil {
		return err
	}
	if len(e) != len(b) {
		return fmt.Errorf("lorawan: exactly %d bytes are expected", len(e))
	}
	copy(e[:], b)
	return nil
}

// String implement fmt.Stringer.
func (e EUI64) String() string {
	return hex.EncodeToString(e[:])
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (e EUI64) MarshalBinary() ([]byte, error) {
	out := make([]byte, len(e))
	// little endian
	for i, v := range e {
		out[len(e)-i-1] = v
	}
	return out, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (e *EUI64) UnmarshalBinary(data []byte) error {
	if len(data) != len(e) {
		return fmt.Errorf("lorawan: %d bytes of data are expected", len(e))
	}
	for i, v := range data {
		// little endian
		e[len(e)-i-1] = v
	}
	return nil
}