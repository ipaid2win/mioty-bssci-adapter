package messages

//go:generate msgp
//msgp:tuple GeoLocation


type GeoLocation struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	Ele float32 `json:"ele"`
}
