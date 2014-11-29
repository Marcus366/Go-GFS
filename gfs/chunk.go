package gfs

import (
	"fmt"
)

type Chunk struct {
	location *ChunkServerMsg
	uuid      uint64
	size      uint64
}

func NewChunk(location *ChunkServerMsg) *Chunk {
	c := new(Chunk)
	c.location = location
	c.size = 0

	return c
}
