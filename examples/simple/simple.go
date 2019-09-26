package main

import (
	"github.com/wetdesertrock/gart"
)

func main() {
	painting := gart.InitProgramWithPainting("simple")

	renderer := painting.Renderer

	curColor := gart.HSLColor{0, 0.5, 0.5}
	curLocation := gart.NewVector2dFromPolar(0, 0.01)
	for i := 0; i < 100; i++ {
		renderer.SetColor(curColor, 0.5)

		renderer.Circle(true, curLocation.X+0.5, curLocation.Y+0.5, 0.01)

		curLocation = curLocation.Rotate(0.15).MultScalar(1.1)
		curColor.H += 0.05
	}

	painting.Save()
}
