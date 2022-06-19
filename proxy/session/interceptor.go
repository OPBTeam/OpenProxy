package session

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

type Listener interface {
	Packet(Source, Player, packet.Packet)
	Leave(Player)
}
