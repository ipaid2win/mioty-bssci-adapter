package bssci_v1

import (
	"bytes"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"
	"mioty-bssci-adapter/internal/common"

	"testing"

	"github.com/tinylib/msgp/msgp"
)

var (
	testMessageConRaw = []byte{77, 73, 79, 84, 89, 66, 48, 49, 160, 0, 0, 0, 137, 167, 99, 111, 109, 109, 97, 110, 100,
		163, 99, 111, 110, 164, 111, 112, 73, 100, 0, 167, 118, 101, 114, 115, 105, 111, 110, 165,
		49, 46, 48, 46, 48, 165, 98, 115, 69, 117, 105, 207, 0, 7, 50, 0, 0, 119, 103, 243, 166,
		118, 101, 110, 100, 111, 114, 174, 68, 105, 101, 104, 108, 32, 77, 101, 116, 101, 114, 105,
		110, 103, 165, 109, 111, 100, 101, 108, 181, 77, 73, 79, 84, 89, 32, 80, 114, 101, 109,
		105, 117, 109, 32, 71, 97, 116, 101, 119, 97, 121, 164, 110, 97, 109, 101, 173, 77, 48, 48,
		48, 55, 51, 50, 55, 55, 54, 55, 70, 51, 164, 98, 105, 100, 105, 195, 168, 115, 110, 66,
		115, 85, 117, 105, 100, 220, 0, 16, 208, 195, 114, 208, 197, 33, 208, 167, 120, 73, 208,
		155, 208, 139, 78, 41, 208, 199, 208, 131, 208, 183, 53, 208, 221}

	testLengthCon  = 160
	testHeaderCon  = testMessageConRaw[0:12]
	testVendor     = "Diehl Metering"
	testModel      = "MIOTY Premium Gateway"
	testName       = "M0007327767F3"
	testMessageCon = messages.Con{
		Command: structs.MsgCon,
		OpId:    0,
		Version: "1.0.0",
		BsEui:   common.EUI64{0x00, 0x07, 0x32, 0x00, 0x00, 0x77, 0x67, 0xF3},
		Vendor:  &testVendor,
		Model:   &testModel,
		Name:    &testName,
		SnBsUuid: [16]int8{-61, 114, -59, 33, -89, 120, 73, -101, -117, 78, 41, -57, -125, -73, 53, -35},
		Bidi:     true,
	}

	testBssciIdentifier = "MIOTYB01"
)

func TestMarshalBssciMessage(t *testing.T) {

	buf, err := MarshalBssciMessage(&testMessageCon)

	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(buf, testMessageConRaw){

		var raw msgp.Raw
		var raw2 msgp.Raw

		raw, _ = testMessageCon.MarshalMsg(nil)
		raw2 = testMessageConRaw[12:]
		json, _ := raw.MarshalJSON()
		json2, _ := raw2.MarshalJSON()
		t.Logf("\n%s\n%s", json, json2)
		t.Errorf("\nexpected \n%v,\ngot bytes for \n%v", testMessageConRaw, buf)
	}
}

func BenchmarkMarshalBssciMessage(b *testing.B) {
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ = MarshalBssciMessage(&testMessageCon)

	}
	_ = buf

}

func TestUnmarshalBssciMessage(t *testing.T) {

	cmd, raw, err := UnmarshalBssciMessage(testMessageConRaw)

	if err != nil {
		t.Fatal(err)
	}

	if cmd != structs.MsgCon {
		t.Errorf("expected %s, got for  %s", structs.MsgCon, cmd)
	}
	json, _ := raw.MarshalJSON()
	t.Logf("\n%s", json)
}

func BenchmarkUnmarshalBssciMessage(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = UnmarshalBssciMessage(testMessageConRaw)

	}
}

func TestCodecBssciMessage(t *testing.T) {

	cmd, raw, err := UnmarshalBssciMessage(testMessageConRaw)

	if err != nil {
		t.Fatal(err)
	}

	if cmd != structs.MsgCon {
		t.Errorf("expected %s, got for  %s", structs.MsgCon, cmd)
	}

	json, _ := raw.MarshalJSON()
	t.Logf("%s", json)

	var con messages.Con
	_, err = con.UnmarshalMsg(raw)
	if err != nil {
		t.Fatal(err)
	}

}

func BenchmarkCodecBssciMessage(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, raw, _ := UnmarshalBssciMessage(testMessageConRaw)
		var con messages.Con
		_, _ = con.UnmarshalMsg(raw)
	}
}

func TestBssciIdentifier(t *testing.T) {

	bssciIdentifierFromBytes := string(bssciIdentifier[:])

	if testBssciIdentifier != bssciIdentifierFromBytes {
		t.Errorf("expected %s, got bytes for  %s", testBssciIdentifier, bssciIdentifierFromBytes)
	}
}

func TestGetBssciMessageLengthFromHeader(t *testing.T) {
	header := []byte{77, 73, 79, 84, 89, 66, 48, 49, 0, 0, 0, 0}
	length, err := getBssciMessageLengthFromHeader(header)

	if err != nil {
		t.Fatal(err)
	}
	if length != 0 {
		t.Errorf("expected %v, got %v", 0, length)
	}
}
func BenchmarkGetBssciMessageLengthFromHeader(b *testing.B) {
	header := []byte{77, 73, 79, 84, 89, 66, 48, 49, 0, 0, 0, 0}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = getBssciMessageLengthFromHeader(header)
	}
}

func TestPrepareBssciMessage(t *testing.T) {
	buf, err := prepareBssciMessage(testLengthCon)

	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(testHeaderCon, buf[0:12]) {
		t.Errorf("expected %b, got bytes for  %b", testHeaderCon, buf)
	}
}

func BenchmarkPrepareBssciMessage(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = prepareBssciMessage(testLengthCon)
	}
}
