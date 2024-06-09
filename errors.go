package ddpgo

import "errors"

var (
	DisconnectErr        = errors.New("Connection closed")
	NoValidSocketAddrErr = errors.New("No valid socket address")
	ParseErr             = errors.New("Parse error")
	UnknownClientErr     = errors.New("Unknown client")
	InvalidPacketErr     = errors.New("Invalid packet")
	NothingToResolveErr  = errors.New("Nothing to resolve")
	CrossBeamErr         = errors.New("Cross beam error")
)
