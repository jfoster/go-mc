package world

import (
	"fmt"
)

type Coord struct {
	X float64
	Y float64
	Z float64

	Dimension Dimension
}

func NewCoord(x float64, y float64, z float64) Coord {
	return Coord{X: x, Y: y, Z: z, Dimension: Overworld}
}

func (c *Coord) Add(coord Coord) Coord {
	return Coord{c.X + coord.X, c.Y + coord.Y, c.Z + coord.Z, coord.Dimension}
}

func (c *Coord) AddX(x float64) Coord {
	return c.Add(Coord{X: x, Y: 0, Z: 0, Dimension: c.Dimension})
}

func (c *Coord) AddY(y float64) Coord {
	return c.Add(Coord{X: 0, Y: y, Z: 0, Dimension: c.Dimension})
}

func (c *Coord) AddZ(z float64) Coord {
	return c.Add(Coord{X: 0, Y: 0, Z: z, Dimension: c.Dimension})
}

func (c *Coord) Nether() Coord {
	if c.Dimension == Nether {
		return *c
	}
	return Coord{c.X / 8, c.Y, c.Z / 8, Nether}
}

func (c Coord) String() string {
	return fmt.Sprintf("{X:%.0f Y:%.0f Z:%.0f}", c.X, c.Y, c.Z)
}

type Coords []Coord

func (c Coords) Len() int      { return len(c) }
func (c Coords) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Coords) Less(i, j int) bool {
	if c[i].X < c[j].X {
		return true
	}
	if c[i].X > c[j].X {
		return false
	}
	return c[i].Z < c[j].Z
}
