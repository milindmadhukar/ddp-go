package protocol

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
