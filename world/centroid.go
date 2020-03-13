package world

import (
	"math"
)

type Centroid struct {
	Coord
	Coords Coords
}

func CentroidOfCoords(coords ...Coord) (c Centroid) {
	c.Coords = coords

	var xsum, ysum, zsum float64

	for _, v := range c.Coords {
		xsum += v.X
		ysum += v.Y
		zsum += v.Z
	}

	len := float64(len(c.Coords))

	xcent := math.Round(xsum / len)
	ycent := math.Round(ysum / len)
	zcent := math.Round(zsum / len)

	c.Coord = Coord{xcent, ycent, zcent, coords[0].Dimension}

	return c
}

func (c *Centroid) Distances() (d []float64) {
	for _, v := range c.Coords {
		dist := math.Sqrt((math.Pow(c.X-v.X, 2) + math.Pow(c.Y-v.Y, 2) + math.Pow(c.Z-v.Z, 2)))
		d = append(d, dist)
	}
	return d
}
