package gart

// Colors brought in from here
// http://www.easyrgb.com/en/math.php

import (
	"math"
)

// HSLColor is a struct used to hold and operate on HSL colors
type HSLColor struct {
	H float64
	S float64
	L float64
}

func NewHSLColor(h, s, l float64) HSLColor {
	return HSLColor{
		H: h,
		S: s,
		L: l,
	}
}

func NewHSLColorFromRGB(r, g, b float64) HSLColor {
	h, s, l := RGBtoHSL(r, g, b)
	return HSLColor{
		H: h,
		S: s,
		L: l,
	}
}

func (this HSLColor) ToRGB() (float64, float64, float64) {
	return HSLtoRGB(this.H, this.S, this.L)
}

func RGBtoHSL(r, g, b float64) (float64, float64, float64) {
	var h, s, l float64

	min := math.Min(r, math.Min(g, b))
	max := math.Max(r, math.Max(g, b))

	deltaMax := max - min

	l = (max + min) / 2

	if deltaMax == 0 {
		h = 0
		s = 0
	} else {
		if l < 0.5 {
			s = deltaMax / (max + min)
		} else {
			s = deltaMax / (2 - max - min)
		}

		deltaR := (((max - r) / 6) + (deltaMax / 2)) / deltaMax
		deltaG := (((max - g) / 6) + (deltaMax / 2)) / deltaMax
		deltaB := (((max - b) / 6) + (deltaMax / 2)) / deltaMax

		if r == max {
			h = deltaB - deltaG
		} else if g == max {
			h = (1 / 3) + deltaR - deltaB
		} else if b == max {
			h = (2 / 3) + deltaG - deltaR
		}

		if h < 0 {
			h += 1
		}
		if h > 1 {
			h -= 1
		}
	}

	return h, s, l
}

func HSLtoRGB(h, s, l float64) (float64, float64, float64) {
	var r, g, b float64

	if s == 0 {
		return l, l, l
	} else {
		var val1, val2 float64

		if l < 0.5 {
			val2 = l * (1 + s)
		} else {
			val2 = (l + s) - (s * l)
		}

		val1 = 2*l - val2

		r = hueToRGB(val1, val2, h+(1.0/3.0))
		g = hueToRGB(val1, val2, h)
		b = hueToRGB(val1, val2, h-(1.0/3.0))
	}

	return r, g, b
}

func hueToRGB(a, b, h float64) float64 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	if (6 * h) < 1 {
		return a + (b-a)*6*h
	}
	if (2 * h) < 1 {
		return b
	}
	if (3 * h) < 2 {
		return a + (b-a)*((2/3)-h)*6
	}

	return a
}
