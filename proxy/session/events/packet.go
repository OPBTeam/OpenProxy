package events

import (
	"github.com/Suremeo/ProxyEye/proxy/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type PacketEventHandler func(event *Context, player session.Player, source session.Source, pk packet.Packet)
