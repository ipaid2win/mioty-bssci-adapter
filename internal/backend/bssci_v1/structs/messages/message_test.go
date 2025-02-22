package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"
)

func TestMessage_UnMarshalMessagePack(t *testing.T) {
	testVendor := "Test Vendor"
	testModel := "Test Model"
	testBsName := "M0007327767F3"
	testVersion := "1.0.0"
	testBsEui := common.EUI64{0x00, 0x07, 0x32, 0x00, 0x00, 0x77, 0x67, 0xF3}
	testBsSessionUuid := structs.SessionUuid{-61, 114, -59, 33, -89, 120, 73, -101, -117, 78, 41, -57, -125, -73, 53, -35}
	testScSessionUuid := structs.SessionUuid{-61, 114, -59, 33, -89, 120, 73, -101, -117, 78, 41, -57, -125, -73, 53, -35}
	testScEui := common.EUI64{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	testScName := "Test Name"

	type fields struct {
		Raw []byte
	}
	tests := []struct {
		name    string
		msg     Message
		fields  fields
		want    Message
		wantErr bool
	}{
		{
			name: "msgCon",
			msg:  &Con{},
			fields: fields{
				Raw: []byte{137, 167, 99, 111, 109, 109, 97, 110, 100, 163, 99, 111, 110, 164, 111, 112, 73, 100, 0, 167, 118, 101, 114, 115, 105, 111, 110, 165, 49, 46, 48, 46, 48, 165, 98, 115, 69, 117, 105, 203, 67, 28, 200, 0, 1, 221, 159, 204, 166, 118, 101, 110, 100, 111, 114, 171, 84, 101, 115, 116, 32, 86, 101, 110, 100, 111, 114, 165, 109, 111, 100, 101, 108, 170, 84, 101, 115, 116, 32, 77, 111, 100, 101, 108, 164, 110, 97, 109, 101, 173, 77, 48, 48, 48, 55, 51, 50, 55, 55, 54, 55, 70, 51, 164, 98, 105, 100, 105, 195, 168, 115, 110, 66, 115, 85, 117, 105, 100, 220, 0, 16, 208, 195, 114, 208, 197, 33, 208, 167, 120, 73, 208, 155, 208, 139, 78, 41, 208, 199, 208, 131, 208, 183, 53, 208, 221},
			}, want: &Con{Command: structs.MsgCon, OpId: 0,
				Version:  testVersion,
				BsEui:    testBsEui,
				Vendor:   &testVendor,
				Model:    &testModel,
				Name:     &testBsName,
				SnBsUuid: testBsSessionUuid,
				Bidi:     true,
			}, wantErr: false,
		},
		{
			name: "msgConRsp",
			msg:  &ConRsp{},
			fields: fields{
				Raw: []byte{137, 167, 99, 111, 109, 109, 97, 110, 100, 166, 99, 111, 110, 82, 115, 112, 164, 111, 112, 73, 100, 0, 167, 118, 101, 114, 115, 105, 111, 110, 165, 49, 46, 48, 46, 48, 165, 98, 115, 69, 117, 105, 203, 67, 112, 16, 16, 16, 16, 16, 16, 166, 118, 101, 110, 100, 111, 114, 171, 84, 101, 115, 116, 32, 86, 101, 110, 100, 111, 114, 165, 109, 111, 100, 101, 108, 170, 84, 101, 115, 116, 32, 77, 111, 100, 101, 108, 164, 110, 97, 109, 101, 169, 84, 101, 115, 116, 32, 78, 97, 109, 101, 168, 115, 110, 82, 101, 115, 117, 109, 101, 194, 168, 115, 110, 83, 99, 85, 117, 105, 100, 220, 0, 16, 208, 195, 114, 208, 197, 33, 208, 167, 120, 73, 208, 155, 208, 139, 78, 41, 208, 199, 208, 131, 208, 183, 53, 208, 221},
			}, want: &ConRsp{
				Command:  structs.MsgConRsp,
				OpId:     0,
				Version:  testVersion,
				ScEui:    testScEui,
				Vendor:   &testVendor,
				Model:    &testModel,
				Name:     &testScName,
				SnResume: false,
				SnScUuid: testScSessionUuid,
			}, wantErr: false,
		},
		{
			name: "msgConCmp",
			msg:  &ConCmp{},
			fields: fields{
				Raw: []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 99, 111, 110, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			}, want: &ConCmp{
				Command: structs.MsgConCmp,
				OpId:    0,
			}, wantErr: false,
		},
		{
			name: "msgPing",
			msg:  &Ping{},
			fields: fields{
				Raw: []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			}, want: &Ping{
				Command: structs.MsgPing,
				OpId:    0,
			}, wantErr: false,
		},
		{
			name: "msgPingRsp",
			msg:  &PingRsp{},
			fields: fields{
				Raw: []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			}, want: &PingRsp{
				Command: structs.MsgPingRsp,
				OpId:    0,
			}, wantErr: false,
		},
		{
			name: "msgPingCmp",
			msg:  &PingCmp{},
			fields: fields{
				Raw: []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			}, want: &PingCmp{
				Command: structs.MsgPingCmp,
				OpId:    0,
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.msg
			_, err := msg.UnmarshalMsg(tt.fields.Raw)

			if (err != nil) != tt.wantErr {
				t.Errorf("Message.UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(msg, tt.want) {
				t.Errorf("Message.UnmarshalMsg() = %v, want %v", msg, tt.want)
			}

		})
	}
}

// func TestMarshalConMessage(t *testing.T) {

// 	buf, err := testMessageCon.MarshalMsg(nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if bytes.Equal(buf, testMessageConRaw) {

// 		var raw msgp.Raw
// 		var raw2 msgp.Raw

// 		raw, _ = testMessageCon.MarshalMsg(nil)
// 		raw2 = testMessageConRaw[12:]
// 		json, _ := raw.MarshalJSON()
// 		json2, _ := raw2.MarshalJSON()
// 		t.Logf("\n%s\n%s", json, json2)
// 		t.Errorf("\nexpected \n%v,\ngot bytes for \n%v", testMessageConRaw, buf)
// 	}
// }

// func TestUnmarshalConMessage(t *testing.T) {

// 	var con Con
// 	_, err := con.UnmarshalMsg(testMessageConRaw)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("\n%v", con)

// }
