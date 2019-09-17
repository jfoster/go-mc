package main

import (
	"fmt"

	mc "github.com/jfoster/mccoordtools"
)

func main() {
	coords := mc.Coords{
		mc.Coord{-20936.0, 32.0, -20938.0},
		mc.Coord{-21136.0, 32.0, -20992.0},
	}

	centroid := mc.CentroidOfCoords(coords...)

	fmt.Println("Overworld:", centroid, "Nether:", centroid.Nether())

	centroid.Coord.Y = 96

	fmt.Println(centroid.Distances())

	fmt.Println("N:", centroid.AddZ(-128))
	fmt.Println("E:", centroid.AddX(128))
	fmt.Println("S:", centroid.AddZ(128))
	fmt.Println("W:", centroid.AddX(-128))

	circle := mc.Circle(centroid.Coord, 257)
	for i, v := range circle {
		i++
		fmt.Println(i, v)
	}
}
