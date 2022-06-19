package world

import (
	"sync"

	"github.com/Suremeo/ProxyEye/proxy/world/chunk"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type World struct {
	name   string
	mutex  sync.Mutex
	store  *Store
	chunks map[chunk.Pos]*Chunk
}

func newWorld(store *Store, name string) *World {
	return &World{
		name:   name,
		mutex:  sync.Mutex{},
		store:  store,
		chunks: make(map[chunk.Pos]*Chunk),
	}
}

func (world *World) GetChunk(pos chunk.Pos) (*Chunk, bool) {
	world.mutex.Lock()
	c, ok := world.chunks[pos]
	world.mutex.Unlock()
	return c, ok
}

func (world *World) AddChunk(pk *packet.LevelChunk) (*Chunk, error) {
	pos := chunk.Pos{pk.Position.X(), pk.Position.Z()}
	ch, err := chunk.NetworkDecode(pk.RawPayload, pk.SubChunkCount)
	if err != nil {
		return nil, err
	}
	world.mutex.Lock()
	c := &Chunk{
		viewers: make(map[uint64]struct{}),
		world:   world,
		Pos:     pos,
		Chunk:   ch,
	}
	world.chunks[pos] = c
	world.mutex.Unlock()
	return c, err
}

func (world *World) removeChunk(pos chunk.Pos) {
	world.mutex.Lock()
	delete(world.chunks, pos)
	world.mutex.Unlock()
}
