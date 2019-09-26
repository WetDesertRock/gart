package gart

import (
	"fmt"

	"github.com/wetdesertrock/flexiconfig"
)

// World is an interface defining a world that will be simulated.
type World interface {
	Init(simulator *Simulator)
	Update(simulator *Simulator, dt float64)
	Render(simulator *Simulator)
}

// Simulator is a struct that provides an easy way of drawing things in a
// simulated fashion similar to game engines. It allows you to define a
// framerate and how long it will run for.
type Simulator struct {
	*Painting
	world World
}

// NewSimulator creates a new simulator struct using the settings object. You
// shouldn't need to ever use this yourself for normal use. Instead you should
// use InitProgramWithSimulator.
func NewSimulator(world World, settings flexiconfig.Settings) *Simulator {
	result := Simulator{
		Painting: NewPainting(settings),
		world:    world,
	}

	return &result
}

// Run will start simulating the world.
func (this *Simulator) Run() {
	totaltime, _ := this.Settings.GetInt("gart:simulator:TotalTime", 0)
	frametime, _ := this.Settings.GetFloat("gart:simulator:FrameTime", 1)

	iterations := float64(totaltime) / frametime

	this.world.Init(this)

	for i := 0.0; i < iterations; i++ {
		this.world.Update(this, frametime)
		this.world.Render(this)
	}

	fmt.Println("done!")
}
