package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"
)

func TestNewDetPrp(t *testing.T) {
	type args struct {
		opId  int64
		epEui common.EUI64
	}
	tests := []struct {
		name string
		args args
		want DetPrp
	}{
		{
			name: "detPrp",
			args: args{1, common.EUI64{1}},
			want: DetPrp{
				Command: structs.MsgDetPrp,
				OpId:    1,
				EpEui:   common.EUI64{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetPrp(tt.args.opId, tt.args.epEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetPrp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "detPrp", fields: fields{structs.MsgDetPrp, 1, common.EUI64{1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetPrp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "detPrp", fields: fields{structs.MsgDetPrp, 1, common.EUI64{1}}, want: structs.MsgDetPrp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetPrp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetPrpRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DetPrpRsp
	}{
		{
			name: "detPrpRsp", args: args{1},
			want: DetPrpRsp{
				Command: structs.MsgDetPrpRsp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetPrpRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetPrpRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "detPrpRsp", fields: fields{structs.MsgDetPrpRsp, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetPrpRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "detPrpRsp", fields: fields{structs.MsgDetPrpRsp, 1}, want: structs.MsgDetPrpRsp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetPrpRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetPrpCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DetPrpCmp
	}{
		{
			name: "detPrpCmp", args: args{1},
			want: DetPrpCmp{
				Command: structs.MsgDetPrpCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetPrpCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetPrpCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpCmp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "detPrpCmp", fields: fields{structs.MsgDetPrpCmp, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetPrpCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpCmp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "detPrpCmp", fields: fields{structs.MsgDetPrpCmp, 1}, want: structs.MsgDetPrpCmp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetPrpCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
