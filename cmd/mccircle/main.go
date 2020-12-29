package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jfoster/go-minecraft/world"
)

func main() {
	var radius = flag.Float64("r", 0, "radius")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Incorrect number of arguments!")
		return
	}

	var coords world.Coords
	var x, y, z float64

	for _, arg := range args {

		split := strings.Split(arg, " ")

		for i, v := range split {
			if i%3 == 0 {
				x, _ = strconv.ParseFloat(v, 64)
			}
			if i%3 == 1 {
				y, _ = strconv.ParseFloat(v, 64)
			}
			if i%3 == 2 {
				z, _ = strconv.ParseFloat(v, 64)

				coord, err := world.NewCoord(x, y, z)
				if err != nil {
					log.Fatal(err)
				}
				coords = append(coords, *coord)
			}
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
