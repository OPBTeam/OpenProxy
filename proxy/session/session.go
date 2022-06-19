package session

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

type Session interface {
	Destination
	Source
	Server() Server
	Packet(source Source, destination Destination, pk packet.Packet) error
	Close()
	SetConnected()
	IsConnected() bool
}
