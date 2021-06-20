package color

import (
	"github.com/cckuailong/colorsys-go"
)

func clampFactor(amount float64) float64 {
	if amount > 1 {
		return 1
	}
	if amount < 0 {
		return 0
	}
	return amount
}

func clampScalar(amount float64) float64 {
	if amount < 0 {
		return 0
	}
	return amount
}

// Brighten a color by some amount
func (c Color) Brighten(amount float64) (o Color) {
	amount = clampScalar(amount)
	h, s, v := colorsys.Rgb2Hsv(uint32(c.R), uint32(c.G), uint32(c.B))
	old := v
	v *= 1 + amount
	// clamp to white
	if v < old {
		v = 1
	}
	r, g, b := colorsys.Hsv2Rgb(h, s, v)
	o.R = uint8(r >> 24)
	o.G = uint8(g >> 24)
	o.B = uint8(b >> 24)
	o.A = c.A
	return
}

// Fade a color to black by some amount
func (c Color) Fade(amount float64) (o Color) {
	amount = clampFactor(amount)
	h, s, v := colorsys.Rgb2Hsv(uint32(c.R), uint32(c.G), uint32(c.B))
	v *= amount
	r, g, b := colorsys.Hsv2Rgb(h, s, v)
	o.R = uint8(r >> 24)
	o.G = uint8(g >> 24)
	o.B = uint8(b >> 24)
	o.A = c.A
	return
}

// Transparent makes a color transparent by a given degree
func (c Color) Transparent(amount float64) (o Color) {
	amount = clampFactor(amount)
	o.A = uint8(float64(c.A) * amount)
	return
}

// Opaque makes a color more opaque by a given degree
func (c Color) Opaque(amount float64) (o Color) {
	amount = clampScalar(amount)
	old := c.A
	o.A = uint8(float64(c.A) * (1 + amount))
	if o.A < old {
		o.A = 0xff
	}
	return
}

// Saturate increases the saturation of a color by a given degree
func (c Color) Saturate(amount float64) (o Color) {
	amount = clampScalar(amount)
	h, s, v := colorsys.Rgb2Hsv(uint32(c.R), uint32(c.G), uint32(c.B))
	s *= 1 + amount
	r, g, b := colorsys.Hsv2Rgb(h, s, v)
	o.R = uint8(r >> 24)
	o.G = uint8(g >> 24)
	o.B = uint8(b >> 24)
	o.A = c.A
	return
}

//Desaturate greys out a color by a given degree.
func (c Color) Desaturate(amount float64) (o Color) {
	amount = clampFactor(amount)
	h, s, v := colorsys.Rgb2Hsv(uint32(c.R), uint32(c.G), uint32(c.B))
	s *= amount
	r, g, b := colorsys.Hsv2Rgb(h, s, v)
	o.R = uint8(r >> 24)
	o.G = uint8(g >> 24)
	o.B = uint8(b >> 24)
	o.A = c.A
	return
}
