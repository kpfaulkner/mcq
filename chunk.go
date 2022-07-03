package mcq

import (
	"github.com/alaingilbert/mcq/mc"
	"github.com/alaingilbert/mcq/nbt"
	"math/bits"
)

const NbSection int = 16
const SectionHeight int = 16
const XDim int = 16
const YDim int = 256
const ZDim int = 16

// Chunk ...
type Chunk struct {
	localX, localZ int
	data           *nbt.NbtTree
}

// NewChunk ...
func NewChunk(localX, localZ int) *Chunk {
	chunk := new(Chunk)
	chunk.localX = localX
	chunk.localZ = localZ
	return chunk
}

func (c *Chunk) GetX() int {
	return c.localX
}

func (c *Chunk) GetZ() int {
	return c.localZ
}

func (c *Chunk) GetData() (data *nbt.NbtTree) {
	return c.data
}

// SetData ...
func (c *Chunk) SetData(data *nbt.NbtTree) {
	c.data = data
}

// Each iterates all blocks in a chunk
func (c *Chunk) Each(clb func(blockID mc.ID, x, y, z int)) {
	sections := c.GetData().Root().Entries["sections"].(*nbt.TagNodeList)
	for s := 0; s < NbSection; s++ {
		section := sections.Get(s).(*nbt.TagNodeCompound)
		blockStates := section.Entries["block_states"].(*nbt.TagNodeCompound)
		palette := blockStates.Entries["palette"].(*nbt.TagNodeList)
		if palette.Length() == 1 {
			continue
		}
		data := blockStates.Entries["data"].(*nbt.TagNodeLongArray)
		mask := uint8(0b1111)
		if palette.Length() > 64 {
			mask = 0b111_1111
		} else if palette.Length() > 32 {
			mask = 0b11_1111
		} else if palette.Length() > 16 {
			mask = 0b1_1111
		}
		ones := bits.OnesCount8(mask)
		for blockPos := 0; blockPos < XDim*ZDim*SectionHeight; blockPos++ {
			blockLngIdx := blockPos / (64 / ones)
			lng := data.Data()[blockLngIdx]
			indexRemaining := blockPos % (64 / ones)
			blockPaletteIndex := int(uint8(lng>>(indexRemaining*ones)) & mask)
			blockID := mc.ID(palette.Get(blockPaletteIndex).(*nbt.TagNodeCompound).Entries["Name"].(*nbt.TagNodeString).String())
			y := blockPos / YDim
			z := blockPos % YDim / XDim
			x := blockPos % XDim
			y += s * SectionHeight
			y -= 64
			clb(blockID, x, y, z)
		}
	}
}
