package main

import (
	"fmt"
	"os"
	"strconv"

	mc "github.com/jfoster/mccoordtool"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args)%3 != 0 {
		fmt.Println("Incorrect number of arguments!")
		return
	}

	var coords mc.Coords
	var x, y, z float64

	for i, v := range args {
		if i%3 == 0 {
			x, _ = strconv.ParseFloat(v, 64)
		}
		if i%3 == 1 {
			y, _ = strconv.ParseFloat(v, 64)
		}
		if i%3 == 2 {
			z, _ = strconv.ParseFloat(v, 64)
			coords = append(coords, mc.Coord{X: x, Y: y, Z: z})
		}
	}

	centroid := mc.CentroidOfCoords(coords...)

	fmt.Println("Overworld:", centroid, "Nether:", centroid.Nether())
	fmt.Println("Distances:", centroid.Distances())

	var radius float64 = 128

	fmt.Println("N:", centroid.AddZ(-radius))
	fmt.Println("E:", centroid.AddX(radius))
	fmt.Println("S:", centroid.AddZ(radius))
	fmt.Println("W:", centroid.AddX(-radius))

	circle := mc.Circle(centroid.Coord, radius*2)
	for i, v := range circle {
		fmt.Println(i+1, v)
	}
}
