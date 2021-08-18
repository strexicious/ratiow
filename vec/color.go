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

func (c Color) WriteColor(w io.Writer) {
	c = c.Scale(255.999)
	fmt.Fprintf(w, "%d %d %d\n", int(c.X()), int(c.Y()), int(c.Z()))
}
