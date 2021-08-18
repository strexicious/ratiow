package vec

import (
	"fmt"
	"io"
)

type Color = Vec3

var ZeroColor = ZeroVec3

var NewColor = NewVec3

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
	c = c.Unscale(float64(samples_per_pixel)).ClampScalar(0, 0.999).Scale(256)
	fmt.Fprintf(w, "%d %d %d\n", int(c.R()), int(c.G()), int(c.B()))
}
