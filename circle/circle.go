package circle

import "github.com/jfoster/mcutils/coord"

func Circle(center coord.Coord, diameter float64) (coords coord.Coords) {
	if even(int64(diameter)) {
		diameter += 1
	}
	for i := 0; i < 4; i++ {
		coords = append(coords, quad(i, center, float64(diameter))...)
	}
	return
}

func quad(q int, center coord.Coord, d float64) (coords coord.Coords) {
	var r = d / 2

	var x, y int64 = int64(center.X), int64(center.Z)

	var x1, y1 int64 = int64(-r), 0

	var e = 2 - d

	for {
		switch q {
		case 0:
			coords = append(coords, coord.Coord{float64(x + y1), center.Y, float64(y + x1), center.Dimension})
		case 1:
			coords = append(coords, coord.Coord{float64(x - x1), center.Y, float64(y + y1), center.Dimension})
		case 2:
			coords = append(coords, coord.Coord{float64(x - y1), center.Y, float64(y - x1), center.Dimension})
		case 3:
			coords = append(coords, coord.Coord{float64(x + x1), center.Y, float64(y - y1), center.Dimension})
		}

		r = e
		if r > float64(x1) {
			x1++
			e += float64(x1*2 + 1)
		}
		if r <= float64(y1) {
			y1++
			e += float64(y1*2 + 1)
		}
		if x1 >= 0 {
			break
		}
	}
	return
}

func even(number int64) bool {
	return number%2 == 0
}
