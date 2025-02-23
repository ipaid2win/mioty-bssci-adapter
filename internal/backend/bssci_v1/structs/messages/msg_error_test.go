package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"reflect"
	"testing"
)

func TestNewBssciError(t *testing.T) {
	type args struct {
		opId    int64
		code    uint32
		message string
	}
	tests := []struct {
		name string
		args args
		want BssciError
	}{
		{
			name: "error",
			args: args{1, 2, "error message"},
			want: BssciError{
				Command: structs.MsgError,
				OpId:    1,
				Code:    2,
				Message: "error message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBssciError(tt.args.opId, tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBssciError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBssciError_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		Code    uint32
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "error",
			fields: fields{Command: structs.MsgError,
				OpId:    1,
				Code:    2,
				Message: "error message",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BssciError{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("BssciError.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBssciError_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		Code    uint32
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "error",
			fields: fields{
				Command: structs.MsgError,
				OpId:    1,
				Code:    2,
				Message: "error message",
			},
			want: structs.MsgError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BssciError{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BssciError.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBssciErrorAck(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want BssciErrorAck
	}{
		{
			name: "errorAck",
			args: args{1},
			want: BssciErrorAck{
				Command: structs.MsgErrorAck,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBssciErrorAck(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBssciErrorAck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBssciErrorAck_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "errorAck",
			fields: fields{
				Command: structs.MsgErrorAck,
				OpId:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BssciErrorAck{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("BssciErrorAck.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBssciErrorAck_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "errorAck",
			fields: fields{
				Command: structs.MsgErrorAck,
				OpId:    1,
			},
			want: structs.MsgErrorAck,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &BssciErrorAck{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BssciErrorAck.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
