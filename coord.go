package mccoordtools

import "fmt"

type Coord struct {
	X float64
	Y float64
	Z float64
}

func (c Coord) String() string {
	return fmt.Sprintf("{X:%.0f Y:%.0f Z:%.0f}", c.X, c.Y, c.Z)
}

func (c Coord) Nether() Coord {
	return Coord{c.X / 8, c.Y, c.Z / 8}
}

func (c Coord) Add(coord Coord) Coord {
	return Coord{c.X + coord.X, c.Y + coord.Y, c.Z + coord.Z}
}

func (c Coord) AddX(x float64) Coord {
	return c.Add(Coord{x, 0, 0})
}

func (c Coord) AddZ(z float64) Coord {
	return c.Add(Coord{0, 0, z})
}

type Coords []Coord

func (c Coords) Len() int      { return len(c) }
func (c Coords) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Coords) Less(i, j int) bool {
	if c[i].X < c[j].X {
		return true
	}
	return c[i].Z < c[j].Z
}
