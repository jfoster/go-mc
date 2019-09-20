package main

import (
	"fmt"

	mc "github.com/jfoster/mccoordtools"
)

func main() {
	centroid := mc.CentroidOfCoords(
		mc.Coord{X: -20936, Y: 32, Z: -20938},
		mc.Coord{X: -21136, Y: 32, Z: -20992},
	)

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
