package master

import (
	"math/rand"
)


type Chunk struct {
	location *ChunkServer
	uuid     uint64
	size     uint64
}

func NewChunk(location *ChunkServer) *Chunk {
	c := new(Chunk)
	c.location = location
	c.size = 0

	rad := rand.Int63()
	for rad <= 0 {
		rad = rand.Int63()
	}
	c.uuid = uint64(rad)

	return c
}
