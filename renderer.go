package gart

import (
	"math"

	"github.com/ungerik/go-cairo"
	"github.com/wetdesertrock/flexiconfig"
)

// Renderer is a struct that holds the cairo instance and exposes wrapper
// methods using cairo. This is the struct you will use to do the actual
// drawaing.
type Renderer struct {
	surface *cairo.Surface
}

// NewRenderer creates a new renderer from a settings instance. You shouldn't
// need to use this for average use. It will be created when you create a new
// painting or simulator.
func NewRenderer(config flexiconfig.Settings) *Renderer {
	result := Renderer{}

	width, _ := config.GetInt("gart:renderer:Width", 100)
	height, _ := config.GetInt("gart:renderer:Height", width)

	result.surface = cairo.NewSurface(cairo.FORMAT_ARGB32, int(width), int(height))

	backgroundColor := HSLColor{}
	config.Get("gart:renderer:BackgroundColor", &backgroundColor)
	backgroundAlpha, _ := config.GetFloat("gart:renderer:BackgroundAlpha", 1)
	r, g, b := backgroundColor.ToRGB()

	result.surface.Scale(float64(width), float64(height))
	result.surface.SetSourceRGBA(r, g, b, backgroundAlpha)
	result.surface.Rectangle(0, 0, 1, 1)
	result.surface.Fill()

	return &result
}

// GetSurface returns the underlying cairo Surface.
func (this *Renderer) GetSurface() *cairo.Surface {
	return this.surface
}

// SetColor sets the color the renderer will use for future operations
func (this *Renderer) SetColor(color HSLColor, alpha float64) {
	r, g, b := color.ToRGB()
	this.surface.SetSourceRGBA(r, g, b, alpha)
}

// SetBlendMode sets the blend mode for cairo.
func (this *Renderer) SetBlendMode(operator cairo.Operator) {
	this.surface.SetOperator(operator)
}

// Circle will draw a circle located at the x/y locatinon
func (this *Renderer) Circle(filled bool, x, y, radius float64) {
	this.surface.NewSubPath()
	this.surface.Arc(x, y, radius, 0, 2*math.Pi)
	this.surface.ClosePath()

	this.surface.Fill()
}

// LineWidth sets the width of the lines drawn
func (this *Renderer) LineWidth(lineWidth float64) {
	this.surface.SetLineWidth(lineWidth)
}

// Line draws a line defined by the connected points.
func (this *Renderer) Line(connected bool, points []Vector2d) {

	this.surface.MoveTo(points[0].X, points[0].Y)

	for i := 1; i < len(points); i++ {
		this.surface.LineTo(points[i].X, points[i].Y)
	}

	if connected {
		this.surface.LineTo(points[0].X, points[0].Y)
	}

	this.surface.Stroke()
}

// Polygon draws a filled polygon defined by the points.
func (this *Renderer) Polygon(points []Vector2d) {
	this.surface.NewSubPath()
	this.surface.MoveTo(points[0].X, points[0].Y)

	for i := 1; i < len(points); i++ {
		this.surface.LineTo(points[i].X, points[i].Y)
	}

	this.surface.ClosePath()
	this.surface.Fill()
}

// SavePNG saves the surface to a png.
func (this *Renderer) SavePNG(outPath string) {
	this.surface.WriteToPNG(outPath)

	// TODO: Move this Finish elsewhere?
	this.surface.Finish()
}
