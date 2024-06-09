package protocol

func (c PixelConfig) ToByte() byte {
	var dataTypeByte byte
	if c.CustomerDefined {
		dataTypeByte |= 0x80 // C bit
	}

	dataTypeByte |= byte(c.DataType) << 3      // TTT bits
	dataTypeByte |= byte(c.PixelFormat) & 0x07 // SSS bits

	return dataTypeByte
}

func GetPixelConfigFromByte(b byte) PixelConfig {
	return PixelConfig{
		CustomerDefined: (b & 0x80) != 0,
		DataType:        LEDDataType((b >> 3) & 0x07),
		PixelFormat:     LEDPixelFormat(b & 0x07),
	}
}

func NewPixelConfig(dataType LEDDataType, pixelFormat LEDPixelFormat, customerDefined bool) PixelConfig {

	return PixelConfig{
		DataType:        dataType,
		PixelFormat:     pixelFormat,
		CustomerDefined: customerDefined,
	}
}

func DefaultPixelConfig() PixelConfig {
	return PixelConfig{
		DataType:        RGB,
		PixelFormat:     Pixel24Bits, // 8 bits for R, G, B respectively
		CustomerDefined: false,
	}
}
