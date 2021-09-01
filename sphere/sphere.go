package sphere

import (
	"math"

	"github.com/strexicious/ratiow/hittable"
	"github.com/strexicious/ratiow/vec"
)

type Sphere struct {
	center vec.Point3
	radius float64
	mat    hittable.Material
}

func NewSphere(cen vec.Point3, r float64, m hittable.Material) Sphere {
	return Sphere{center: cen, radius: r, mat: m}
}

func (s *Sphere) Hit(r vec.Ray, t_min, t_max float64) (hit bool, rec *hittable.HitRecord) {
	oc := r.Origin().Sub(s.center)
	a := r.Direction().NormSquared()
	half_b := oc.Dot(r.Direction())
	c := oc.NormSquared() - s.radius*s.radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		hit = false
		return
	}
	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-half_b - sqrtd) / a
	if root <= t_min || t_max <= root {
		root = (-half_b + sqrtd) / a
		if root <= t_min || t_max <= root {
			hit = false
			return
		}
	}

	rec = new(hittable.HitRecord)
	rec.T = root
	rec.P = r.At(root)
	rec.Mat = s.mat

	outward_normal := rec.P.Sub(s.center).Unscale(s.radius)
	rec.SetFaceNormal(r, outward_normal)

	hit = true
	return
}
