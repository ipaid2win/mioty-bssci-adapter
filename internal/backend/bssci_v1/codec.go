package bssci_v1

import (
	"bytes"
	"encoding/binary"
	"io"

	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

const (
	bssciHeaderSize             = 12
	bssciHeaderIdentifierOffset = 0
	bssciHeaderIdentifierSize   = 8
	bssciHeaderLengthOffset     = 8
	bssciHeaderLengthSize       = 4
)

var (
	bssciIdentifier = [8]byte{0x4D, 0x49, 0x4F, 0x54, 0x59, 0x42, 0x30, 0x31}
)

func ReadBssciMessage(r io.Reader) (cmd structs.CommandHeader, raw msgp.Raw, err error) {
	// reader the bssci header
	buf := make([]byte, 12)
	_, err = r.Read(buf)

	if err != nil {
		err = errors.Wrap(err, "io read error on header")
		return
	}

	// get the size of the message
	length, err := getBssciMessageLengthFromHeader(buf)
	if err != nil {
		return
	}
	buf = make([]byte, length)
	_, err = r.Read(buf)

	if err != nil {
		err = errors.Wrap(err, "io read error on message")
		return
	}

	// parse out command
	_, err = cmd.UnmarshalMsg(buf)
	if err != nil {
		err = errors.Wrap(err, "message error")
		return
	}

	// get raw message
	_, err = raw.UnmarshalMsg(buf)
	if err != nil {
		err = errors.Wrap(err, "message error")
		return
	}

	return
}

func WriteBssciMessage(w io.Writer, msg messages.Message) (err error) {
	// marshal message
	buf, err := MarshalBssciMessage(msg)
	if err != nil {
		return
	}

	// write message
	_, err = w.Write(buf)

	if err != nil {
		err = errors.Wrap(err, "io write error")
		return
	}

	return
}


// convert msg to message pack and attach bssci header
func MarshalBssciMessage(msg messages.Message) (buf []byte, err error) {
	msgBuf, err := msg.MarshalMsg(nil)
	if err != nil {
		err = errors.Wrap(err, "bssci message marshal error")
		return
	}

	msgLength := len(msgBuf)

	buf, err = prepareBssciMessage(msgLength)

	if err != nil {
		err = errors.Wrap(err, "message header error")
		return
	}
	// add message pack data
	buf = append(buf, msgBuf...)
	return
}

// read the bssci header and extract the raw message pack data
func UnmarshalBssciMessage(buf []byte) (cmd structs.Command, raw msgp.Raw, err error) {

	// get length
	length, err := getBssciMessageLengthFromHeader(buf)
	if err != nil {
		err = errors.Wrap(err, "message error")
		return
	}

	// slice off header
	buf = buf[bssciHeaderSize : bssciHeaderSize+length]

	// parse out command
	var commandHeader structs.CommandHeader
	_, err = commandHeader.UnmarshalMsg(buf)
	if err != nil {
		err = errors.Wrap(err, "message error")
		return
	}
	cmd = commandHeader.GetCommand()

	_, err = raw.UnmarshalMsg(buf)
	if err != nil {
		err = errors.Wrap(err, "message error")
		return
	}
	return
}

// build header and allocate buffer
func prepareBssciMessage(length int) (buf []byte, err error) {
	// allocate buf
	buf = make([]byte, 0, bssciHeaderSize+length)
	// add identifier to header
	buf = append(buf, bssciIdentifier[:]...)
	// add length to header
	buf, err = binary.Append(buf, binary.LittleEndian, uint32(length))
	if err != nil {
		err = errors.Wrap(err, "message header error")
		return
	}
	return
}

func getBssciMessageLengthFromHeader(buf []byte) (l int, err error) {

	if len(buf) < bssciHeaderSize {
		err = errors.New("message header error: invalid header size")
		return
	}

	identifier := [8]byte(buf[bssciHeaderIdentifierOffset : bssciHeaderIdentifierOffset+bssciHeaderIdentifierSize])
	if identifier != bssciIdentifier {
		err = errors.New("message header error: invalid identifier: " + string(bssciIdentifier[:]))
		return
	}

	length := [4]byte(buf[bssciHeaderLengthOffset : bssciHeaderLengthOffset+bssciHeaderLengthSize])
	err = binary.Read(bytes.NewReader(length[:]), binary.LittleEndian, &l)
	if err != nil {
		err = errors.Wrap(err, "message header error")
		return
	}
	return
}
