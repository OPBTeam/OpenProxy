package world

import (
	"bytes"
	"sync"

	"github.com/opbteam/proxyeye/proxy/world/chunk"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Chunk struct {
	*chunk.Chunk
	mutex   sync.Mutex
	world   *World
	Pos     chunk.Pos
	viewers map[uint64]struct{}
}

func (c *Chunk) AddViewer(id uint64) {
	c.mutex.Lock()
	c.viewers[id] = struct{}{}
	c.mutex.Unlock()
}

func (c *Chunk) RemoveViewer(id uint64) {
	c.mutex.Lock()
	delete(c.viewers, id)
	if len(c.viewers) == 0 {
		c.world.removeChunk(c.Pos)
	}
	c.mutex.Unlock()
}

func (c *Chunk) SendTo(conn *minecraft.Conn) error {
	var chunkBuf bytes.Buffer
	data := chunk.NetworkEncode(c.Chunk)

	count := byte(0)
	for y := byte(0); y < 16; y++ {
		if data.SubChunks[y] != nil {
			count = y + 1
		}
	}
	for y := byte(0); y < count; y++ {
		if data.SubChunks[y] == nil {
			_ = chunkBuf.WriteByte(chunk.SubChunkVersion)
			// We write zero here, meaning the sub chunk has no block storages: The sub chunk is completely
			// empty.
			_ = chunkBuf.WriteByte(0)
			continue
		}
		_, _ = chunkBuf.Write(data.SubChunks[y])
	}
	_, _ = chunkBuf.Write(data.Data2D)
	_, _ = chunkBuf.Write(data.BlockNBT)

	return conn.WritePacket(&packet.LevelChunk{
		BlobHashes:      []uint64{0},
		CacheEnabled:    false,
		Position:        protocol.ChunkPos(c.Pos),
		SubChunkCount:   uint32(count),
		HighestSubChunk: uint16(count),
		RawPayload:      append([]byte(nil), chunkBuf.Bytes()...),
	})
}
