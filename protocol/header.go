package protocol

func (h DDPHeader) ToBytes() DDPHeaderBytes {
	var b DDPHeaderBytes

	b[0] = h.PacketConfig.ToByte()
	b[1] = h.SequenceNumber
	b[2] = h.PixelConfig.ToByte()
	b[3] = GetSourceDestinationIDByte(h.SourceDestinationID)

	b[4] = byte(h.Offset >> 24)
	b[5] = byte(h.Offset >> 16)
	b[6] = byte(h.Offset >> 8)
	b[7] = byte(h.Offset)

	b[8] = byte(h.Length >> 8)
	b[9] = byte(h.Length)

	return b
}

// TODO: To implement when timecode is implemented
func (h DDPHeader) ToBytesTimecode() DDPHeaderBytesTimecode {
	var b DDPHeaderBytesTimecode

	return b
}

func GetDDPHeaderFromBytes(b DDPHeaderBytes) DDPHeader {
	return DDPHeader{
		PacketConfig:        GetPacketConfigFromByte(b[0]),
		SequenceNumber:      b[1],
		PixelConfig:         GetPixelConfigFromByte(b[2]),
		SourceDestinationID: GetSourceDestinationIDFromByte(b[3]),
		Offset:              uint32(b[4])<<24 | uint32(b[5])<<16 | uint32(b[6])<<8 | uint32(b[7]),
		Length:              uint16(b[8])<<8 | uint16(b[9]),
	}
}

// TODO: To implement when timecode is implemented
func GetDDPHeaderFromBytesTimecode(b DDPHeaderBytesTimecode) DDPHeader {
  return DDPHeader{}
}

func DefaultDDPHeader() DDPHeader {
  return DDPHeader{
    PacketConfig:        DefaultPacketConfig(),
    SequenceNumber:      0,
    PixelConfig:         DefaultPixelConfig(),
    SourceDestinationID: DefaultID,
    Offset:              0,
    Length:              0,
  }
}
