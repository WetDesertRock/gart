package main

import (
	// "fmt"
	"container/list"
	"math"
	"math/rand"

	"github.com/wetdesertrock/gart"
)

type Paint struct {
	points *list.List
	center gart.Vector2d
	color  gart.HSLColor
}

func NewPaint(center gart.Vector2d, radius float64) *Paint {
	paint := Paint{}
	paint.points = list.New()

	paint.center = center

	paint.color = gart.HSLColor{H: 0.2, S: 0.5, L: 0.5}

	maxpoints := 20.0
	angle := math.Pi * 2 / maxpoints
	for i := 0.0; i < maxpoints; i++ {
		theta := angle * i
		point := gart.NewVector2dFromPolar(theta, radius).Add(center)

		paint.points.PushBack(point)
	}

	return &paint
}

func subdivide(a, b gart.Vector2d) gart.Vector2d {
	midpoint := a.Sub(b).MultScalar(0.5).Add(b)

	curlength := a.Distance(b)

	mag := (rand.NormFloat64()*0.5 + 0.5) * curlength * 0.3
	angle := rand.Float64() * math.Pi * 2

	return midpoint.Add(gart.NewVector2dFromPolar(angle, mag))
}

func (this *Paint) Deform(iterations int) *list.List {
	points := list.New()

	// Clone list
	for e := this.points.Front(); e != nil; e = e.Next() {
		point := e.Value.(gart.Vector2d)

		points.PushBack(point)
	}

	for i := 0; i < iterations; i++ {
		start_e := points.Front()
		previous := start_e.Value.(gart.Vector2d)

		for e := start_e.Next(); e != nil; e = e.Next() {
			point := e.Value.(gart.Vector2d)

			points.InsertBefore(subdivide(previous, point), e)
			previous = point
		}

		end_e := points.Back()
		start := start_e.Value.(gart.Vector2d)
		end := end_e.Value.(gart.Vector2d)

		points.PushBack(subdivide(start, end))
	}

	return points
}

func (this *Paint) Update(dt float64) {}

func (this *Paint) Draw(sim *gart.Simulator) {
	renderer := sim.Renderer
	renderer.SetColor(this.color, 0.05)
	renderer.LineWidth(0.005)

	// Get deformation
	points := this.Deform(8)

	polygon := make([]gart.Vector2d, 0, points.Len())

	for e := points.Front(); e != nil; e = e.Next() {
		point := e.Value.(gart.Vector2d)

		polygon = append(polygon, point)
	}

	renderer.Polygon(polygon)
}

type PainterlyWorld struct {
	paints *list.List
}

func NewPainterlyWorld() *PainterlyWorld {
	world := PainterlyWorld{}
	world.paints = list.New()

	return &world
}

func (world *PainterlyWorld) Init(sim *gart.Simulator) {
	world.paints.PushBack(NewPaint(gart.NewVector2d(0.5, 0.5), 0.1))
}

func (world *PainterlyWorld) Update(sim *gart.Simulator, dt float64) {
	// Iterate through list and print its contents.
	for e := world.paints.Front(); e != nil; e = e.Next() {
		paint := e.Value.(*Paint)

		paint.Update(dt)
	}
}

func (world *PainterlyWorld) Render(sim *gart.Simulator) {
	for e := world.paints.Front(); e != nil; e = e.Next() {
		paint := e.Value.(*Paint)

		paint.Draw(sim)
	}
}

func main() {
	sim := gart.InitProgramWithSimulator("painterly", NewPainterlyWorld())
	sim.Run()
	sim.Save()
}
