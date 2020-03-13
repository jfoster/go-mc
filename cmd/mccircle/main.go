package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/jfoster/go-minecraft/world"
)

func main() {
	var radius = flag.Float64("r", 0, "radius")

	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	if len(args) == 0 || len(args)%3 != 0 {
		fmt.Println("Incorrect number of arguments!")
		return
	}

	var coords world.Coords
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
			coords = append(coords, world.Coord{X: x, Y: y, Z: z})
		}
	}

	cent := world.CentroidOfCoords(coords...)

	fmt.Println("Overworld:", cent.String(), "Nether:", cent.Nether().String())
	fmt.Println("Distances:", cent.Distances())

	if r := *radius; r > 0 {
		fmt.Println("N:", cent.AddZ(-r))
		fmt.Println("E:", cent.AddX(r))
		fmt.Println("S:", cent.AddZ(r))
		fmt.Println("W:", cent.AddX(-r))

		circle := world.Circle(cent.Coord, r*2)
		for i, v := range circle {
			fmt.Println(i+1, v)
		}
	}
}
