package protocol

func (c PacketConfig) ToByte() byte {
	var b byte
	b |= flagVersion1 // NOTE: Default Version

	if c.Timecode {
		b |= flagTimecode
	}
	if c.Storage {
		b |= flagStorage
	}
	if c.Reply {
		b |= flagReply
	}
	if c.Query {
		b |= flagQuery
	}
	if c.Push {
		b |= flagPush
	}
	return b
}

func GetPacketConfigFromByte(b byte) PacketConfig {
	c := PacketConfig{}
	c.Version = flagVersion1 // NOTE: Default Version

	if b&flagTimecode != 0 {
		c.Timecode = true
	}
	if b&flagStorage != 0 {
		c.Storage = true
	}
	if b&flagReply != 0 {
		c.Reply = true
	}
	if b&flagQuery != 0 {
		c.Query = true
	}
	if b&flagPush != 0 {
		c.Push = true
	}
	return c
}

func DefaultPacketConfig() PacketConfig {
	return PacketConfig{
		Version:  flagVersion1,
		Timecode: false,
		Storage:  false,
		Reply:    false,
		Query:    false,
    Push:     false, // NOTE: Maybe needs to be true
	}
}
