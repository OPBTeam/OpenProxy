package session

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

type Destination interface {
	WritePacket(packet.Packet) error
}
