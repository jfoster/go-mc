package centroid

import (
	"math"

	"github.com/jfoster/mcutils/coord"
)

type Centroid struct {
	coord.Coord
	coord.Coords
}

func CentroidOfCoords(coords ...coord.Coord) (c Centroid) {
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

	c.Coord = coord.Coord{xcent, ycent, zcent, coords[0].Dimension}

	return c
}

func (c *Centroid) Distances() (d []float64) {
	for _, v := range c.Coords {
		dist := math.Sqrt((math.Pow(c.Coord.X-v.X, 2) + math.Pow(c.Coord.Y-v.Y, 2) + math.Pow(c.Coord.Z-v.Z, 2)))
		d = append(d, dist)
	}
	return d
}
