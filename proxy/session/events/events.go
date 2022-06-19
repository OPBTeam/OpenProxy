package events

type Event int

const (
	EventAntiCheatDetection = iota
	EventConnect
	EventJoin
	EventLeave
	EventPacket
)
