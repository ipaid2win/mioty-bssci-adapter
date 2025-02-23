package messages

import (
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewStatus(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want Status
	}{
		{name: "status", args: args{1}, want: Status{
			Command: structs.MsgStatus,
			OpId:    1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatus(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "status", fields: fields{structs.MsgStatus, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Status{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("Status.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "status", fields: fields{structs.MsgStatus, 1}, want: structs.MsgStatus},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Status{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Status.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatusRsp(t *testing.T) {
	type args struct {
		opId        int64
		code        uint32
		message     string
		time        uint64
		dutyCycle   float32
		geoLocation *GeoLocation
		uptime      *uint64
		temp        *float64
		cpuLoad     *float64
		memLoad     *float64
	}
	tests := []struct {
		name string
		args args
		want StatusRsp
	}{
		{name: "statusRsp", args: args{1, 0, "test", 1000, 0.5, nil, nil, nil, nil, nil}, want: StatusRsp{
			Command:     structs.MsgStatusRsp,
			OpId:        1,
			Code:        0,
			Message:     "test",
			Time:        1000,
			DutyCycle:   0.5,
			GeoLocation: nil,
			Uptime:      nil,
			Temp:        nil,
			CpuLoad:     nil,
			MemLoad:     nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusRsp(tt.args.opId, tt.args.code, tt.args.message, tt.args.time, tt.args.dutyCycle, tt.args.geoLocation, tt.args.uptime, tt.args.temp, tt.args.cpuLoad, tt.args.memLoad); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		Code        uint32
		Message     string
		Time        uint64
		DutyCycle   float32
		GeoLocation *GeoLocation
		Uptime      *uint64
		Temp        *float64
		CpuLoad     *float64
		MemLoad     *float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "statusRsp", fields: fields{
			Command:     structs.MsgStatusRsp,
			OpId:        1,
			Code:        0,
			Message:     "",
			Time:        0,
			DutyCycle:   0,
			GeoLocation: &GeoLocation{},
			Uptime:      new(uint64),
			Temp:        new(float64),
			CpuLoad:     new(float64),
			MemLoad:     new(float64),
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &StatusRsp{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				Code:        tt.fields.Code,
				Message:     tt.fields.Message,
				Time:        tt.fields.Time,
				DutyCycle:   tt.fields.DutyCycle,
				GeoLocation: tt.fields.GeoLocation,
				Uptime:      tt.fields.Uptime,
				Temp:        tt.fields.Temp,
				CpuLoad:     tt.fields.CpuLoad,
				MemLoad:     tt.fields.MemLoad,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("StatusRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		Code        uint32
		Message     string
		Time        uint64
		DutyCycle   float32
		GeoLocation *GeoLocation
		Uptime      *uint64
		Temp        *float64
		CpuLoad     *float64
		MemLoad     *float64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "statusRsp", fields: fields{
			Command:     structs.MsgStatusRsp,
			OpId:        1,
			Code:        0,
			Message:     "",
			Time:        0,
			DutyCycle:   0,
			GeoLocation: &GeoLocation{},
			Uptime:      new(uint64),
			Temp:        new(float64),
			CpuLoad:     new(float64),
			MemLoad:     new(float64),
		}, want: structs.MsgStatusRsp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &StatusRsp{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				Code:        tt.fields.Code,
				Message:     tt.fields.Message,
				Time:        tt.fields.Time,
				DutyCycle:   tt.fields.DutyCycle,
				GeoLocation: tt.fields.GeoLocation,
				Uptime:      tt.fields.Uptime,
				Temp:        tt.fields.Temp,
				CpuLoad:     tt.fields.CpuLoad,
				MemLoad:     tt.fields.MemLoad,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusRsp_IntoProto(t *testing.T) {

	var testTime uint64 = 1000000005

	var testUptime uint64 = 1000
	var testTemp float64 = 45.5
	var testCpu float64 = 0.5
	var testMemory float64 = 0.6

	testTs := timestamppb.Timestamp{
		Seconds: int64(1000000),
		Nanos:   int32(5),
	}

	type fields struct {
		Command     structs.Command
		OpId        int64
		Code        uint32
		Message     string
		Time        uint64
		DutyCycle   float32
		GeoLocation *GeoLocation
		Uptime      *uint64
		Temp        *float64
		CpuLoad     *float64
		MemLoad     *float64
	}
	type args struct {
		bsEui common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *msg.BasestationStatus
	}{
		{
			name: "statusRsp1",
			fields: fields{
				Command:     structs.MsgStatusRsp,
				OpId:        1,
				Code:        0,
				Message:     "test",
				Time:        testTime,
				DutyCycle:   0.5,
				GeoLocation: nil,
				Uptime:      nil,
				Temp:        nil,
				CpuLoad:     nil,
				MemLoad:     nil,
			},
			args: args{common.EUI64{1}},
			want: &msg.BasestationStatus{
				BsEui:       1,
				StatusCode:  0,
				StatusMsg:   "test",
				Ts:          &testTs,
				DutyCycle:   0.5,
				GeoLocation: nil,
				Uptime:      nil,
				Temp:        nil,
				Cpu:         nil,
				Memory:      nil,
			},
		},
		{
			name: "statusRsp2",
			fields: fields{
				Command:     structs.MsgStatusRsp,
				OpId:        1,
				Code:        0,
				Message:     "test",
				Time:        testTime,
				DutyCycle:   0.5,
				GeoLocation: &GeoLocation{},
				Uptime:      nil,
				Temp:        nil,
				CpuLoad:     nil,
				MemLoad:     nil,
			},
			args: args{common.EUI64{1}},
			want: &msg.BasestationStatus{
				BsEui:       1,
				StatusCode:  0,
				StatusMsg:   "test",
				Ts:          &testTs,
				DutyCycle:   0.5,
				GeoLocation: &msg.GeoLocation{},
				Uptime:      nil,
				Temp:        nil,
				Cpu:         nil,
				Memory:      nil,
			},
		},

		{
			name: "statusRsp3",
			fields: fields{
				Command:     structs.MsgStatusRsp,
				OpId:        1,
				Code:        0,
				Message:     "test",
				Time:        testTime,
				DutyCycle:   0.5,
				GeoLocation: &GeoLocation{},
				Uptime:      &testUptime,
				Temp:        &testTemp,
				CpuLoad:     &testCpu,
				MemLoad:     &testMemory,
			},
			args: args{common.EUI64{1}},
			want: &msg.BasestationStatus{
				BsEui:       1,
				StatusCode:  0,
				StatusMsg:   "test",
				Ts:          &testTs,
				DutyCycle:   0.5,
				GeoLocation: &msg.GeoLocation{},
				Uptime:      &testUptime,
				Temp:        &testTemp,
				Cpu:         &testCpu,
				Memory:      &testMemory,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &StatusRsp{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				Code:        tt.fields.Code,
				Message:     tt.fields.Message,
				Time:        tt.fields.Time,
				DutyCycle:   tt.fields.DutyCycle,
				GeoLocation: tt.fields.GeoLocation,
				Uptime:      tt.fields.Uptime,
				Temp:        tt.fields.Temp,
				CpuLoad:     tt.fields.CpuLoad,
				MemLoad:     tt.fields.MemLoad,
			}
			got := m.IntoProto(tt.args.bsEui)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusRsp.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatusCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want StatusCmp
	}{
		{name: "statusCmp", args: args{1}, want: StatusCmp{
			Command: structs.MsgStatusCmp,
			OpId:    1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusCmp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "statusCmp", fields: fields{structs.MsgStatusCmp, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &StatusCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("StatusCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusCmp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "statusCmp", fields: fields{structs.MsgStatusCmp, 1}, want: structs.MsgStatusCmp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &StatusCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
