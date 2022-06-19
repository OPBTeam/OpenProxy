package session

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func init() {
	packet.Register(packet.IDCraftingData, func() packet.Packet { return &packet.Unknown{PacketID: packet.IDCraftingData} })
	packet.Register(packet.IDCreativeContent, func() packet.Packet { return &packet.Unknown{PacketID: packet.IDCreativeContent} })
}
