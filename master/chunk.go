package master

import ()

type Chunk struct {
	location *ChunkServer
	uuid     uint64
	size     uint64
}

func NewChunk(location *ChunkServer) *Chunk {
	c := new(Chunk)
	c.location = location
	c.size = 0

	return c
}
