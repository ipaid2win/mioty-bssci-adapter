package messages

import (
	"encoding/json"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestMessageData struct {
	name    string
	msgType Message
	raw     []byte
	msg     Message
	wantErr bool
	json    string
}

type TestMessageSuite struct {
	suite.Suite

	data []TestMessageData
}

func TestMessage(t *testing.T) {
	suite.Run(t, new(TestMessageSuite))
}

func (ts *TestMessageSuite) SetupSuite() {
	testVendor := "Test Vendor"
	testModel := "Test Model"
	testVersion := "1.0.0"

	testBsName := "M0007327767F3"
	testScName := "Test Name"

	testBsSessionUuid := structs.SessionUuid{-61, 114, -59, 33, -89, 120, 73, -101, -117, 78, 41, -57, -125, -73, 53, -35}
	testScSessionUuid := structs.SessionUuid{-61, 114, -59, 33, -89, 120, 73, -101, -117, 78, 41, -57, -125, -73, 53, -35}

	testBsEui := common.EUI64{0x00, 0x07, 0x32, 0x00, 0x00, 0x77, 0x67, 0xF3}
	testScEui := common.EUI64{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	testEpEui := common.EUI64{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	var testStatusUptime uint64 = 1000
	var testStatusTemp float64 = 45.5
	var testStatusCpu float64 = 0.5
	var testStatusMemory float64 = 0.6

	testCon := TestMessageData{
		name:    "msgCon",
		msgType: &Con{},
		raw:     []byte{137, 167, 99, 111, 109, 109, 97, 110, 100, 163, 99, 111, 110, 164, 111, 112, 73, 100, 0, 167, 118, 101, 114, 115, 105, 111, 110, 165, 49, 46, 48, 46, 48, 165, 98, 115, 69, 117, 105, 203, 67, 28, 200, 0, 1, 221, 159, 204, 166, 118, 101, 110, 100, 111, 114, 171, 84, 101, 115, 116, 32, 86, 101, 110, 100, 111, 114, 165, 109, 111, 100, 101, 108, 170, 84, 101, 115, 116, 32, 77, 111, 100, 101, 108, 164, 110, 97, 109, 101, 173, 77, 48, 48, 48, 55, 51, 50, 55, 55, 54, 55, 70, 51, 164, 98, 105, 100, 105, 195, 168, 115, 110, 66, 115, 85, 117, 105, 100, 220, 0, 16, 208, 195, 114, 208, 197, 33, 208, 167, 120, 73, 208, 155, 208, 139, 78, 41, 208, 199, 208, 131, 208, 183, 53, 208, 221},
		msg: &Con{Command: structs.MsgCon, OpId: 0,
			Version:  testVersion,
			BsEui:    testBsEui,
			Vendor:   &testVendor,
			Model:    &testModel,
			Name:     &testBsName,
			SnBsUuid: testBsSessionUuid,
			Bidi:     true,
		},
		wantErr: false,
		json: `{
    "command": "con",
    "opId": 0,
    "version": "1.0.0",
    "bsEui": 2025300426188787,
    "vendor": "Test Vendor",
    "model": "Test Model",
    "name": "M0007327767F3",
    "bidi": true,
    "snBsUuid": [
        -61,
        114,
        -59,
        33,
        -89,
        120,
        73,
        -101,
        -117,
        78,
        41,
        -57,
        -125,
        -73,
        53,
        -35
    ]
}`,
	}

	testConRsp := TestMessageData{
		name:    "msgConRsp",
		msgType: &ConRsp{},
		raw:     []byte{137, 167, 99, 111, 109, 109, 97, 110, 100, 166, 99, 111, 110, 82, 115, 112, 164, 111, 112, 73, 100, 0, 167, 118, 101, 114, 115, 105, 111, 110, 165, 49, 46, 48, 46, 48, 165, 115, 99, 69, 117, 105, 203, 67, 112, 16, 16, 16, 16, 16, 16, 166, 118, 101, 110, 100, 111, 114, 171, 84, 101, 115, 116, 32, 86, 101, 110, 100, 111, 114, 165, 109, 111, 100, 101, 108, 170, 84, 101, 115, 116, 32, 77, 111, 100, 101, 108, 164, 110, 97, 109, 101, 169, 84, 101, 115, 116, 32, 78, 97, 109, 101, 168, 115, 110, 82, 101, 115, 117, 109, 101, 194, 168, 115, 110, 83, 99, 85, 117, 105, 100, 220, 0, 16, 208, 195, 114, 208, 197, 33, 208, 167, 120, 73, 208, 155, 208, 139, 78, 41, 208, 199, 208, 131, 208, 183, 53, 208, 221},
		msg: &ConRsp{
			Command:  structs.MsgConRsp,
			OpId:     0,
			Version:  testVersion,
			ScEui:    testScEui,
			Vendor:   &testVendor,
			Model:    &testModel,
			Name:     &testScName,
			SnResume: false,
			SnScUuid: testScSessionUuid,
		},
		wantErr: false,
		json: `{
	"command": "conRsp",
    "opId": 0,
    "version": "1.0.0",
    "scEui": 72340172838076670,
    "vendor": "Test Vendor",
    "model": "Test Model",
    "name": "Test Name",
    "snResume": false,
    "snScUuid": [
        -61,
        114,
        -59,
        33,
        -89,
        120,
        73,
        -101,
        -117,
        78,
        41,
        -57,
        -125,
        -73,
        53,
        -35
    ]
}`,
	}

	testConCmp := TestMessageData{
		name:    "msgConCmp",
		msgType: &ConCmp{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 99, 111, 110, 67, 109, 112, 164, 111, 112, 73, 100, 0},
		msg: &ConCmp{
			Command: structs.MsgConCmp,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "conCmp",
	"opId": 0
}`,
	}

	testPing := TestMessageData{
		name:    "msgPing",
		msgType: &Ping{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
		msg: &Ping{
			Command: structs.MsgPing,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "ping",
	"opId": 0
}`,
	}
	testPingRsp := TestMessageData{
		name:    "msgPingRsp",
		msgType: &PingRsp{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 82, 115, 112, 164, 111, 112, 73, 100, 0},
		msg: &PingRsp{
			Command: structs.MsgPingRsp,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "pingRsp",
	"opId": 0
}`,
	}

	testPingCmp := TestMessageData{
		name:    "msgPingCmp",
		msgType: &PingCmp{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 67, 109, 112, 164, 111, 112, 73, 100, 0},
		msg: &PingCmp{
			Command: structs.MsgPingCmp,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "pingCmp",
	"opId": 0
}`,
	}

	testStatus := TestMessageData{
		name:    "msgStatus",
		msgType: &Status{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 115, 116, 97, 116, 117, 115, 164, 111, 112, 73, 100, 0},
		msg: &Status{
			Command: structs.MsgStatus,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "status",
	"opId": 0
}`,
	}

	testStatusRsp := TestMessageData{
		name:    "msgStatusRsp",
		msgType: &StatusRsp{},
		raw:     []byte{138, 167, 99, 111, 109, 109, 97, 110, 100, 169, 115, 116, 97, 116, 117, 115, 82, 115, 112, 164, 111, 112, 73, 100, 0, 164, 99, 111, 100, 101, 0, 167, 109, 101, 115, 115, 97, 103, 101, 162, 111, 107, 164, 116, 105, 109, 101, 206, 59, 154, 202, 5, 169, 100, 117, 116, 121, 67, 121, 99, 108, 101, 202, 62, 204, 204, 205, 166, 117, 112, 116, 105, 109, 101, 205, 3, 232, 164, 116, 101, 109, 112, 203, 64, 70, 192, 0, 0, 0, 0, 0, 167, 99, 112, 117, 76, 111, 97, 100, 203, 63, 224, 0, 0, 0, 0, 0, 0, 167, 109, 101, 109, 76, 111, 97, 100, 203, 63, 227, 51, 51, 51, 51, 51, 51},
		msg: &StatusRsp{
			Command:     structs.MsgStatusRsp,
			OpId:        0,
			Code:        0,
			Message:     "ok",
			Time:        1000000005,
			DutyCycle:   0.4,
			GeoLocation: nil,
			Uptime:      &testStatusUptime,
			Temp:        &testStatusTemp,
			CpuLoad:     &testStatusCpu,
			MemLoad:     &testStatusMemory,
		},
		wantErr: false,
		json: `{
	"command": "statusRsp",
	"opId": 0,
	"code": 0,
	"message": "ok",
	"time": 1000000005,
	"dutyCycle": 0.4,
	"uptime": 1000,
	"temp": 45.5,
	"cpuLoad": 0.5,
	"memLoad": 0.6
}`,
	}

	testStatusCmp := TestMessageData{
		name:    "msgStatusCmp",
		msgType: &StatusCmp{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 115, 116, 97, 116, 117, 115, 67, 109, 112, 164, 111, 112, 73, 100, 0},
		msg: &StatusCmp{
			Command: structs.MsgStatusCmp,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "statusCmp",
	"opId": 0
}`,
	}

	testDetPrp := TestMessageData{
		name:    "msgDetPrp",
		msgType: &DetPrp{},
		raw:     []byte{131, 167, 99, 111, 109, 109, 97, 110, 100, 166, 100, 101, 116, 80, 114, 112, 164, 111, 112, 73, 100, 0, 165, 101, 112, 69, 117, 105, 203, 67, 112, 32, 48, 64, 80, 96, 112},
		msg: &DetPrp{
			Command: structs.MsgDetPrp,
			OpId:    0,
			EpEui: testEpEui,
		},
		wantErr: false,
		json: `{
	"command": "detPrp",
	"opId": 0,
	"epEui": 72623859790382856
}`,
	}

	testDetPrpRsp := TestMessageData{
		name:    "msgDetPrpRsp",
		msgType: &DetPrpRsp{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 100, 101, 116, 80, 114, 112, 82, 115, 112, 164, 111, 112, 73, 100, 0},
		msg: &DetPrpRsp{
			Command: structs.MsgDetPrpRsp,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "detPrpRsp",
	"opId": 0
}`,
	}

	testDetPrpCmp := TestMessageData{
		name:    "msgDetPrpCmp",
		msgType: &DetPrpCmp{},
		raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 100, 101, 116, 80, 114, 112, 67, 109, 112, 164, 111, 112, 73, 100, 0},
		msg: &DetPrpCmp{
			Command: structs.MsgDetPrpCmp,
			OpId:    0,
		},
		wantErr: false,
		json: `{
	"command": "detPrpCmp",
	"opId": 0
}`,
	}

	ts.data = []TestMessageData{testCon, testConRsp, testConCmp, testPing, testPingRsp, testPingCmp, testStatus, testStatusRsp, testStatusCmp, testDetPrp, testDetPrpRsp, testDetPrpCmp}

}

func (ts *TestMessageSuite) TestMessage_UnmarshalMessagePack() {
	t := ts.T()
	assert := assert.New(t)

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.msgType
			_, err := msg.UnmarshalMsg(tt.raw)

			if tt.wantErr {
				if !assert.Error(err, "Message.UnmarshalMsg() expected error = %v, got value: %v", err, msg) {
					t.Errorf("Message.UnmarshalMsg() expected error = %v, got value: %v", err, msg)
				}
			} else {
				if !assert.NoError(err, "Message.UnmarshalMsg() unexpected error = %v", err) {
					t.Errorf("Message.UnmarshalMsg() unexpected error = %v", err)
				} else if !assert.Equal(tt.msg, msg, "Message.UnmarshalMsg() = %v, want %v", msg, tt.msg) {
					t.Errorf("Message.UnmarshalMsg() = %v, want %v", msg, tt.msg)
				}

			}
		})
	}
}

func (ts *TestMessageSuite) TestMessage_MarshalMessagePack() {
	t := ts.T()
	assert := assert.New(t)

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.msg
			raw, err := msg.MarshalMsg(nil)

			if tt.wantErr {
				if !assert.Error(err, "Message.MarshalMsg() expected error = %v, got value: %v", err, raw) {
					t.Errorf("Message.MarshalMsg() expected error = %v, got value: %v", err, raw)
				}
			} else {
				if !assert.NoError(err, "Message.MarshalMsg() unexpected error = %v", err) {
					t.Errorf("Message.MarshalMsg() unexpected error = %v", err)
				} else if !assert.Equal(tt.raw, raw, "Message.MarshalMsg() = %v, want %v", raw, tt.raw) {
					t.Errorf("Message.MarshalMsg() = %v, want %v", raw, tt.raw)
				}
			}

		})
	}
}

func (ts *TestMessageSuite) TestMessage_UnmarshalJson() {
	t := ts.T()
	assert := assert.New(t)

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.msgType
			err := json.Unmarshal([]byte(tt.json), msg)

			if tt.wantErr {
				if !assert.Error(err, "Message.UnmarshalJson() expected error = %v, got value: %v", err, msg) {
					t.Errorf("Message.UnmarshalJson() expected error = %v, got value: %v", err, msg)
				}
			} else {
				if !assert.NoError(err, "Message.UnmarshalJson() unexpected error = %v", err) {
					t.Errorf("Message.UnmarshalJson() unexpected error = %v", err)
				} else if !assert.Equal(tt.msg, msg, "Message.UnmarshalJson() = %v, want %v", msg, tt.msg) {
					t.Errorf("Message.UnmarshalJson() = %v, want %v", msg, tt.msg)
				}

			}
		})
	}
}

func (ts *TestMessageSuite) TestMessage_MarshalJson() {
	t := ts.T()
	assert := assert.New(t)

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.msg

			jsonRaw, err := json.MarshalIndent(msg, "", "\t")

			value := string(jsonRaw)

			if tt.wantErr {
				if !assert.Error(err, "Message.MarshalMsg() expected error = %v, got value: %v", err, value) {
					t.Errorf("Message.MarshalMsg() expected error = %v, got value: %v", err, value)
				}
			} else {
				if !assert.NoError(err, "Message.MarshalMsg() unexpected error = %v", err) {
					t.Errorf("Message.MarshalMsg() unexpected error = %v", err)
				} else if !assert.Equal(tt.json, value, "Message.MarshalMsg() = %v, want %v", value, tt.json) {
					t.Errorf("Message.MarshalMsg() = %v, want %v", value, tt.json)
				}
			}

		})
	}
}
