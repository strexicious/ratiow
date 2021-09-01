package materials

import (
	"github.com/strexicious/ratiow/hittable"
	"github.com/strexicious/ratiow/vec"
)

type Lambertian struct {
	Albedo vec.Color
}

func (m *Lambertian) Scatter(r_in vec.Ray, rec *hittable.HitRecord) (bool, vec.Ray, vec.Color) {
	scattered_dir := rec.Normal.Add(vec.RandomUnitVec3())

	if scattered_dir.IsNearZero() {
		scattered_dir = rec.Normal
	}

	return true, vec.NewRay(rec.P, scattered_dir), m.Albedo
}
