package protocol

import "encoding/json"

type PacketConfig struct {
	Version  uint8 // 0x40
	Timecode bool  // 0x10
	Storage  bool  // 0x08
	Reply    bool  // 0x04
	Query    bool  // 0x02
	Push     bool  // 0x01
}

type LEDDataType uint8

type LEDPixelFormat uint8

type PixelConfig struct {
	DataType        LEDDataType
	PixelFormat     LEDPixelFormat
	CustomerDefined bool
}

type DDPHeader struct {
	PacketConfig        PacketConfig
	SequenceNumber      uint8
	PixelConfig         PixelConfig
	SourceDestinationID SourceDestinationID
	Offset              uint32
	Length              uint16
	Timcode             any // TODO: To implement
}

type DDPHeaderBytes [10]byte
type DDPHeaderBytesTimecode [14]byte

type Status struct {
	Update string `json:"update,omitempty"`
	State  string `json:"state,omitempty"`
	Man    string `json:"man,omitempty"`
	Model  string `json:"mod,omitempty"`
	Ver    string `json:"ver,omitempty"`
	Mac    string `json:"mac,omitempty"`
	Push   bool   `json:"push,omitempty"`
	Ntp    bool   `json:"ntp,omitempty"`
}

type StatusRoot struct {
	Status Status `json:"status"`
}

type Port struct {
	Port uint32 `json:"port"`
	Ts   uint32 `json:"ts"`
	L    uint32 `json:"l"`
	Ss   uint32 `json:"ss"`
}

type Config struct {
	IP    string `json:"ip,omitempty"`
	NM    string `json:"nm,omitempty"`
	GW    string `json:"gw,omitempty"`
	Ports []Port `json:"ports"`
}

type ConfigRoot struct {
	Config Config `json:"config"`
}

type ControlRoot struct {
	Control Control `json:"control"`
}

type Color struct {
	R uint32 `json:"r"`
	G uint32 `json:"g"`
	B uint32 `json:"b"`
}

type Control struct {
	Fx     string  `json:"fx,omitempty"`
	Int    uint32  `json:"int,omitempty"`
	Spd    uint32  `json:"spd,omitempty"`
	Dir    uint32  `json:"dir,omitempty"`
	Colors []Color `json:"colors,omitempty"`
	Save   uint32  `json:"save,omitempty"`
	Power  uint32  `json:"power,omitempty"`
}

type Message struct {
	MessageType string       `json:"message_type"`
	Control     *ControlRoot `json:"control,omitempty"`
	Status      *StatusRoot  `json:"status,omitempty"`
	Config      *ConfigRoot  `json:"config,omitempty"`
	Parsed      *struct {
		ID    uint            `json:"id"`
		Value json.RawMessage `json:"value"`
	} `json:"parsed,omitempty"`
	Unparsed *struct {
		ID    uint            `json:"id"`
		Value json.RawMessage `json:"value"`
	} `json:"unparsed,omitempty"`
}
