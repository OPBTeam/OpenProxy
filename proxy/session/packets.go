package session

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func init() {
	packet.RegisterPacketFromClient(packet.IDCraftingData, func() packet.Packet { return &packet.Unknown{PacketID: packet.IDCraftingData} })
	packet.RegisterPacketFromClient(packet.IDCreativeContent, func() packet.Packet { return &packet.Unknown{PacketID: packet.IDCreativeContent} })

}
