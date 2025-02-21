package messages

import "mioty-bssci-adapter/internal/api/msg"

//go:generate msgp
//msgp:tuple GeoLocation

type GeoLocation struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	Alt float32 `json:"alt"`
}

func (m *GeoLocation) IntoProto() *msg.GeoLocation {
	var message msg.GeoLocation

	if m != nil {
		message = msg.GeoLocation{
			Lat: m.Lat,
			Lon: m.Lon,
			Alt: m.Alt,
		}
	}

	return &message
}
