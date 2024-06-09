package protocol

// Config Flags
const (
	flagVersionMask byte = 0xc0
	flagVersion1    byte = 0x40
	flagPush        byte = 0x01
	flagQuery       byte = 0x02
	flagReply       byte = 0x04
	flagStorage     byte = 0x08
	flagTimecode    byte = 0x10
)

/*
	LED Data Types for TTT

// 000 = undefined
// 001 = RGB
// 010 = HSL
// 011 = RGBW
// 100 = grayscale
*/
const (
	UndefinedType LEDDataType = iota
	RGB
	HSL
	RGBW
	Grayscale
)

/*
LED Pixel Formats for SSS

SSS is size in bits per pixel element (like just R or G or B data)
0=undefined, 1=1, 2=4, 3=8, 4=16, 5=24, 6=32
*/
const (
	UndefinedPixelFormat LEDPixelFormat = iota
	Pixel1Bits
	Pixel4Bits
	Pixel8Bits
	Pixel16Bits
	Pixel24Bits
	Pixel32Bits
)

const MAX_DATA_LEN int = 480 * 3
