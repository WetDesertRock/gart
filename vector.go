package gart

import (
	"fmt"
	"math"
)

// Vector2d is a immutable struct for doing 2d vector operations.
type Vector2d struct {
	X float64
	Y float64
}

// NewVector2d creates a new vector from the x and y components
func NewVector2d(x, y float64) Vector2d {
	return Vector2d{x, y}
}

// NewVector2dFromPolar creates a new vector from a given angle and maagnitude.
func NewVector2dFromPolar(angle, magnitude float64) Vector2d {
	x := math.Cos(angle) * magnitude
	y := math.Sin(angle) * magnitude
	return NewVector2d(x, y)
}

// Addition

func (this Vector2d) Add(other Vector2d) Vector2d {
	return NewVector2d(this.X+other.X, this.Y+other.Y)
}

func (this Vector2d) AddScalar(scalar float64) Vector2d {
	return NewVector2d(this.X+scalar, this.Y+scalar)
}

// Subtraction

func (this Vector2d) Sub(other Vector2d) Vector2d {
	return NewVector2d(this.X-other.X, this.Y-other.Y)
}

func (this Vector2d) SubScalar(scalar float64) Vector2d {
	return NewVector2d(this.X-scalar, this.Y-scalar)
}

// Multiply

func (this Vector2d) Mult(other Vector2d) Vector2d {
	return NewVector2d(this.X*other.X, this.Y*other.Y)
}

func (this Vector2d) MultScalar(scalar float64) Vector2d {
	return NewVector2d(this.X*scalar, this.Y*scalar)
}

// Divide

func (this Vector2d) Div(other Vector2d) Vector2d {
	return NewVector2d(this.X/other.X, this.Y/other.Y)
}

func (this Vector2d) DivScalar(scalar float64) Vector2d {
	return NewVector2d(this.X/scalar, this.Y/scalar)
}

// Misc Math

// TODO: This function
// func (this Vector2d) Cross(other Vector2d) (Vector2d) {
// }

// Length returns the length of the vector.
func (this Vector2d) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y)
}

// LengthSq returns the length of the vector squared. This should provide
// performance benefits when in tight loops by not doing math.Sqrt().
func (this Vector2d) LengthSq() float64 {
	return this.X*this.X + this.Y*this.Y
}

// Distance returns the distance computed between two vectors.
func (this Vector2d) Distance(other Vector2d) float64 {
	dx := this.X - other.X
	dy := this.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Normalize returns a normalized vector.
func (this Vector2d) Normalize() Vector2d {
	len := this.Length()
	return this.DivScalar(len)
}

// Direction gives the angle that this vector is facing (in radians)
func (this Vector2d) Direction() float64 {
	return math.Atan2(this.Y, this.X)
}

// PolarCoordinates returns the angle and magnitude of the vector.
func (this Vector2d) PolarCoordinates() (float64, float64) {
	return this.Direction(), this.Length()
}

func (this Vector2d) Rotate(theta float64) Vector2d {
	x := this.X*math.Cos(theta) - this.Y*math.Sin(theta)
	y := this.X*math.Sin(theta) + this.Y*math.Cos(theta)
	return NewVector2d(x, y)
}

// Compare

func (this Vector2d) Equals(other Vector2d) bool {
	// TODO: Make this use an optional elipson value
	return this.X == other.X && this.Y == other.Y
}

//Misc

func (this Vector2d) String() string {
	return fmt.Sprintf("Vector2d{%.3f, %.3f}", this.X, this.Y)
}
