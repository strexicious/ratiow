package materials

import (
	"math"
	"math/rand"

	"github.com/strexicious/ratiow/hittable"
	"github.com/strexicious/ratiow/vec"
)

type Dielectric struct {
	Ir float64
}

func reflectance(cosine, ref_idx float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - ref_idx) / (1 + ref_idx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

func (m *Dielectric) Scatter(r_in vec.Ray, rec *hittable.HitRecord) (bool, vec.Ray, vec.Color) {
	refraction_ratio := m.Ir
	if rec.FrontFace {
		refraction_ratio = 1 / m.Ir
	}

	unit_dir := r_in.Direction().Normalised()

	cos_theta := math.Min(-unit_dir.Dot(rec.Normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta*cos_theta)

	cannot_refract := refraction_ratio*sin_theta > 1

	var direction vec.Vec3

	if cannot_refract || reflectance(cos_theta, refraction_ratio) > rand.Float64() {
		direction = unit_dir.Reflect(rec.Normal)
	} else {
		direction = unit_dir.Refract(rec.Normal, refraction_ratio)
	}

	return true, vec.NewRay(rec.P, direction), vec.NewColor(1, 1, 1)
}
