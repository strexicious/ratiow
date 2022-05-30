package vec

import (
	"fmt"
	"io"
	"math"

	"github.com/strexicious/ratiow/utils"
)

type Color = Vec3

var ZeroColor = ZeroVec3
var NewColor = NewVec3
var RandomColor = RandomVec3
var RandomColorRange = RandomVec3Range

func (c Color) R() float64 {
	return c.e[0]
}

func (c Color) G() float64 {
	return c.e[1]
}

func (c Color) B() float64 {
	return c.e[2]
}

func (c Color) WriteColor(w io.Writer, samples_per_pixel int32) {
	c = c.Unscale(float64(samples_per_pixel))
	r, g, b := math.Sqrt(c.R()), math.Sqrt(c.G()), math.Sqrt(c.B())

	r = 256 * utils.Clamp(r, 0, 0.999)
	g = 256 * utils.Clamp(g, 0, 0.999)
	b = 256 * utils.Clamp(b, 0, 0.999)

	fmt.Fprintf(w, "%d %d %d\n", int(r), int(g), int(b))
}
