package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"reflect"
	"testing"
)

func TestNewPing(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want Ping
	}{
		{
			name: "ping",
			args: args{1},
			want: Ping{
				Command: structs.MsgPing,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPing(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPing_GetOpId(t *testing.T) {
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
			name: "ping",
			fields: fields{
				structs.MsgPingRsp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Ping{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("Ping.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPing_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "ping", fields: fields{structs.MsgPing, 1}, want: structs.MsgPing},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Ping{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ping.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPingRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want PingRsp
	}{
		{name: "pingRsp", args: args{1}, want: PingRsp{
			Command: structs.MsgPingRsp,
			OpId:    1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPingRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPingRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPingRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "pingRsp", fields: fields{structs.MsgPingRsp, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PingRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("PingRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPingRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "pingRsp", fields: fields{structs.MsgPingRsp, 1}, want: structs.MsgPingRsp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PingRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PingRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPingCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want PingCmp
	}{
		{name: "pingCmp", args: args{1}, want: PingCmp{
			Command: structs.MsgPingCmp,
			OpId:    1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPingCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPingCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPingCmp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "pingCmp", fields: fields{structs.MsgPingCmp, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PingCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("PingCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPingCmp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "pingCmp", fields: fields{structs.MsgPingCmp, 1}, want: structs.MsgPingCmp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PingCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PingCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
