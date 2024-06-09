package ddpgo

import (
	"net"

	"github.com/milindmadhukar/ddp-go/protocol"
)

type DDPConnection struct {
	PixelConfig         protocol.PixelConfig
	SourceDestinationID protocol.SourceDestinationID
	sequenceNumber      uint8
	udpConnection       *net.UDPConn
	address             *net.UDPAddr

	Reciever chan []byte
	buffer   []byte
}

func DefaultDDPConnection(addr string, port int) (*DDPConnection, error) {
	ddpConn := DDPConnection{
		PixelConfig:         protocol.DefaultPixelConfig(),
		SourceDestinationID: protocol.DefaultID,
		sequenceNumber:      1,
		address:             &net.UDPAddr{IP: net.ParseIP(addr), Port: port},
		Reciever:            make(chan []byte),
		buffer:              make([]byte, 1500),
	}

	udpConn, err := net.DialUDP("udp", nil, ddpConn.address)
	if err != nil {
		return nil, err
	}

	ddpConn.udpConnection = udpConn

	return &ddpConn, nil
}

func (conn *DDPConnection) Close() error {
	return conn.udpConnection.Close()
}

func (conn *DDPConnection) Write(data []byte) (int, error) {
	header := protocol.DefaultDDPHeader()
	header.PacketConfig.Push = false
	header.PixelConfig = conn.PixelConfig

	return conn.SendData(&header, data)
}

func (conn *DDPConnection) WriteOffset(data []byte, offset uint32) (int, error) {
	header := protocol.DefaultDDPHeader()
	header.PacketConfig.Push = false
	header.PixelConfig = conn.PixelConfig
	header.Offset = offset

	return conn.SendData(&header, data)
}

func (conn *DDPConnection) assemblePacket(header *protocol.DDPHeader, data []byte) int {

	if header.PacketConfig.Timecode {
		panic("Timecode not implemented")
		// use header.ToBytesTimecode()
	} else {
		headerBytes := header.ToBytes()
		copy(conn.buffer, headerBytes[:10])

		copy(conn.buffer[10:], data)
		return len(headerBytes) + len(data)
	}
}

func (conn *DDPConnection) SendData(header *protocol.DDPHeader, data []byte) (int, error) {
	offset := int(header.Offset)
	sent := 0

	num_iterations := (len(data) + protocol.MAX_DATA_LEN - 1) / protocol.MAX_DATA_LEN
	iter := 0

	for offset < len(data) {
		iter++

		if iter == num_iterations {
			header.PacketConfig.Push = true
		}

		header.SequenceNumber = conn.sequenceNumber

		chunkEnd := min(offset+protocol.MAX_DATA_LEN, len(data))
		chunk := data[offset:chunkEnd]
		header.Length = uint16(len(chunk))

		// NOTE: To be implemented when timecode is implemented to check if packet is 10 bytes or 14 bytes
		packetLen := conn.assemblePacket(header, chunk)

		currentSent, err := conn.udpConnection.Write(conn.buffer[:packetLen])
		if err != nil {
			return sent, err
		}

		sent += currentSent

		if conn.sequenceNumber > 15 {
			conn.sequenceNumber = 1
		} else {
			conn.sequenceNumber++
		}

		offset += protocol.MAX_DATA_LEN
		header.Offset = uint32(offset)
	}

	return sent, nil
}
