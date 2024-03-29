package session

import (
	"github.com/sandertv/gophertunnel/minecraft"
)

type Player interface {
	Source
	Destination
	Chunks() Chunks
	IsTransferring() bool
	SetTransferring(bool)
	Session() Session
	SetSession(Session)
	Raknet() *minecraft.Conn
	Anticheat() AntiCheatProfile

	Online() bool
	Chat(msg string, format ...interface{})
	Kick(...string)
	Sound(name string, volume, pitch float32)

	Close()
}
