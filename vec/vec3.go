package vec

import (
	"fmt"
	"io"
	"math"
	"math/rand"

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

func (v Vec3) IsNearZero() bool {
	EPSILON := 1e-8
	return math.Abs(v.e[0]) < EPSILON &&
		math.Abs(v.e[1]) < EPSILON &&
		math.Abs(v.e[2]) < EPSILON
}

func (v Vec3) Reflect(normal Vec3) Vec3 {
	return v.Sub(normal.Scale(2 * v.Dot(normal)))
}

func (v Vec3) Refract(normal Vec3, eta_ratio float64) Vec3 {
	cos_theta := math.Min(-v.Dot(normal), 1.0)
	r_out_perp := normal.Scale(cos_theta).Add(v).Scale(eta_ratio)
	r_out_parallel := normal.Scale(-math.Sqrt(math.Abs(1 - r_out_perp.NormSquared())))
	return r_out_perp.Add(r_out_parallel)
}

func (v Vec3) ComponentWiseScale(s Vec3) Vec3 {
	v.e[0] *= s.e[0]
	v.e[1] *= s.e[1]
	v.e[2] *= s.e[2]
	return v
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

func RandomVec3() Vec3 {
	x, y, z := rand.Float64(), rand.Float64(), rand.Float64()
	return NewVec3(x, y, z)
}

func RandomVec3Range(min, max float64) Vec3 {
	return NewVec3(min, min, min).Add(RandomVec3().Scale(max - min))
}

func RandomVec3InsideUnitSphere() Vec3 {
	for {
		v := RandomVec3Range(-1, 1)
		if v.NormSquared() < 1 {
			return v
		}
	}
}

func RandomVec3InsideHemisphere(normal Vec3) Vec3 {
	random := RandomVec3InsideUnitSphere()
	if random.Dot(normal) > 0 {
		return random
	} else {
		return random.Neg()
	}
}

func RandomUnitVec3() Vec3 {
	return RandomVec3InsideUnitSphere().Normalised()
}

func RandomInUnitDisk() Vec3 {
	for {
		p := NewVec3(rand.Float64()*2-1, rand.Float64()*2-1, 0)
		if p.NormSquared() >= 1 {
			continue
		}
		return p
	}
}
