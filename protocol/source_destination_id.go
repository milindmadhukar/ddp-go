package protocol

// TODO: weird.

type SourceDestinationID uint8

const (
	ReservedID SourceDestinationID = 0
	DefaultID  SourceDestinationID = 1
	// Custom from 2 to 246
	ControlID   SourceDestinationID = 249
	ConfigID    SourceDestinationID = 250
	StatusID    SourceDestinationID = 251
	DMXID       SourceDestinationID = 254
	BroadcastID SourceDestinationID = 255
)

func GetSourceDestinationIDByte(id SourceDestinationID) byte {
	return byte(id)
}

func GetSourceDestinationIDFromByte(b byte) SourceDestinationID {
  return SourceDestinationID(b)
}
