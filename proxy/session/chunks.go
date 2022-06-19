package session

import (
	"github.com/Suremeo/ProxyEye/proxy/world"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Chunks interface {
	Player() Player
	Move(vec3 mgl32.Vec3)
	World() *world.World
	SetWorld(*world.World)
	Chunk() (chunk *world.Chunk, ok bool)
	AddChunk(pk *packet.LevelChunk)
	GetBlock(pos protocol.BlockPos) uint32
	Floor(vec3 mgl32.Vec3) uint32
	UpdateBlock(pos protocol.BlockPos, block uint32)
	Radius() int
	UpdateRadius(int)
	Close()
}
