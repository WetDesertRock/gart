package main

import (
	// "fmt"
	"container/list"
	"math"
	"math/rand"

	"github.com/wetdesertrock/gart"
)

type Particle struct {
	Position gart.Vector2d
	Velocity gart.Vector2d
	Life     float64
	Lifespan float64
	Rotation float64
	Color    gart.HSLColor
}

func (this *Particle) Update(dt float64) {
	this.Position = this.Position.Add(this.Velocity.MultScalar(dt))
	this.Velocity = this.Velocity.Rotate(this.Rotation * dt)
	this.Life -= dt
}

type TestWorld struct {
	particles *list.List
}

func NewTestWorld() *TestWorld {
	world := TestWorld{}
	world.particles = list.New()

	return &world
}

func (world *TestWorld) Init(sim *gart.Simulator) {
	for i := 0; i < 1000; i++ {
		dir := rand.Float64() * math.Pi * 2
		mag := rand.Float64()*0.1 + 0.1
		lifespan := rand.Float64()*0.4 + 3
		h := (rand.Float64()-0.5)*0.1 + 0.9
		s := (rand.Float64()-0.5)*0.1 + 0.7
		l := float64(0.5)

		particle := Particle{
			Position: gart.NewVector2d(0.5, 0.5),
			Velocity: gart.NewVector2dFromPolar(dir, mag),
			Life:     lifespan,
			Lifespan: lifespan,
			Rotation: rand.Float64()*0.3 - 0.15,
			Color:    gart.HSLColor{H: h, S: s, L: l},
		}

		world.particles.PushBack(&particle)
	}
}

func (world *TestWorld) Update(sim *gart.Simulator, dt float64) {
	for e := world.particles.Front(); e != nil; e = e.Next() {
		particle := e.Value.(*Particle)

		particle.Update(dt)

		if particle.Life <= 0 {
			world.particles.Remove(e)
		}
	}
}

func (world *TestWorld) Render(sim *gart.Simulator) {
	renderer := sim.Renderer

	for e := world.particles.Front(); e != nil; e = e.Next() {
		particle := e.Value.(*Particle)

		lifefactor := particle.Life / particle.Lifespan

		renderer.SetColor(particle.Color, 0.7*lifefactor)

		renderer.Circle(true, particle.Position.X, particle.Position.Y, 0.001)
	}
}

func main() {
	sim := gart.InitProgramWithSimulator("particles", NewTestWorld())
	sim.Run()
	sim.Save()
}
