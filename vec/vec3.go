package vec

import (
	"fmt"
	"io"
	"math"

	"github.com/strexicious/ratiow/utils"
)

type Vec3 struct {
	e [3]float64
}

func ZeroVec3() Vec3 {
	return Vec3{e: [3]float64{0, 0, 0}}
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{e: [3]float64{x, y, z}}
}

func (v Vec3) X() float64 {
	return v.e[0]
}

func (v Vec3) Y() float64 {
	return v.e[1]
}

func (v Vec3) Z() float64 {
	return v.e[2]
}

func (v Vec3) Neg() Vec3 {
	v.e[0] = -v.e[0]
	v.e[1] = -v.e[1]
	v.e[2] = -v.e[2]
	return v
}

func (v *Vec3) Ith(i int) *float64 {
	return &v.e[i]
}

func (v Vec3) Add(u Vec3) Vec3 {
	v.e[0] += u.e[0]
	v.e[1] += u.e[1]
	v.e[2] += u.e[2]
	return v
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return v.Add(u.Neg())
}

func (v Vec3) Scale(c float64) Vec3 {
	v.e[0] *= c
	v.e[1] *= c
	v.e[2] *= c
	return v
}

func (v Vec3) Unscale(c float64) Vec3 {
	v.e[0] /= c
	v.e[1] /= c
	v.e[2] /= c
	return v
}

func (v Vec3) NormSquared() float64 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

func (v Vec3) Norm() float64 {
	return math.Sqrt(v.NormSquared())
}

func (v Vec3) Length() float64 {
	return v.Norm()
}

func (v Vec3) Dot(u Vec3) float64 {
	return v.X()*u.X() + v.Y()*u.Y() + v.Z()*u.Z()
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return NewVec3(
		v.Y()*u.Z()-v.Z()*u.Y(),
		v.Z()*u.X()-v.X()*u.Z(),
		v.X()*u.Y()-v.Y()*u.X(),
	)
}

func (v Vec3) Normalised() Vec3 {
	return v.Unscale(v.Norm())
}

func (v Vec3) ClampScalar(min, max float64) Vec3 {
	v.e[0] = utils.Clamp(v.e[0], min, max)
	v.e[1] = utils.Clamp(v.e[1], min, max)
	v.e[2] = utils.Clamp(v.e[2], min, max)
	return v
}

func (v Vec3) WriteVec3(w io.Writer) {
	fmt.Fprintf(w, "%g %g %g\n", v.X(), v.Y(), v.Z())
}
