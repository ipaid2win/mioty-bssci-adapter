package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"reflect"
	"testing"
)

func TestMessage_UnMarshalMessagePack(t *testing.T) {
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
		{name: "msgCon", msg: &Con{},
			fields: fields{
				Raw: []byte{137, 167, 99, 111, 109, 109, 97, 110, 100,
					163, 99, 111, 110, 164, 111, 112, 73, 100, 0, 167, 118, 101, 114, 115, 105, 111, 110, 165,
					49, 46, 48, 46, 48, 165, 98, 115, 69, 117, 105, 207, 0, 7, 50, 0, 0, 119, 103, 243, 166,
					118, 101, 110, 100, 111, 114, 174, 68, 105, 101, 104, 108, 32, 77, 101, 116, 101, 114, 105,
					110, 103, 165, 109, 111, 100, 101, 108, 181, 77, 73, 79, 84, 89, 32, 80, 114, 101, 109,
					105, 117, 109, 32, 71, 97, 116, 101, 119, 97, 121, 164, 110, 97, 109, 101, 173, 77, 48, 48,
					48, 55, 51, 50, 55, 55, 54, 55, 70, 51, 164, 98, 105, 100, 105, 195, 168, 115, 110, 66,
					115, 85, 117, 105, 100, 220, 0, 16, 208, 195, 114, 208, 197, 33, 208, 167, 120, 73, 208,
					155, 208, 139, 78, 41, 208, 199, 208, 131, 208, 183, 53, 208, 221},
			}, want: &Con{Command: structs.MsgCon, OpId: 0,
				Version:  "1.0.0",
				BsEui:    2025300426188787,
				Vendor:   &testVendor,
				Model:    &testModel,
				Name:     &testName,
				SnBsUuid: [16]int8{-61, 114, -59, 33, -89, 120, 73, -101, -117, 78, 41, -57, -125, -73, 53, -35},
				Bidi:     true,
			}, wantErr: false,
		}}
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
