package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesCircle   sides = 0
	SidesTriangle sides = 3
	SidesSquare   sides = 4
)

type sides int

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	if sidesNum == 4 {
		fmt.Println(sideLen * sideLen)
	} else if sidesNum == 3 {
		fmt.Println((3 / 4) * sideLen * sideLen)
	} else if sidesNum == 0 {
		fmt.Println(math.Pi * sideLen)
	} else {
		fmt.Println("it wasn't our figure")
	}
	return 1.0
}

func main() {
	CalcSquare(10.0, SidesCircle)
	CalcSquare(10.0, SidesSquare)
	CalcSquare(10.0, SidesTriangle)

}
