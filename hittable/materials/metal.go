package materials

import (
	"github.com/strexicious/ratiow/hittable"
	"github.com/strexicious/ratiow/vec"
)

type Metal struct {
	Albedo vec.Color
	Fuzz   float64
}

func (m *Metal) Scatter(r_in vec.Ray, rec *hittable.HitRecord) (bool, vec.Ray, vec.Color) {
	reflected := r_in.Direction().Normalised().Reflect(rec.Normal)
	scattered := vec.NewRay(rec.P, reflected.Add(vec.RandomVec3InsideUnitSphere().Scale(m.Fuzz)))

	if reflected.Dot(rec.Normal) > 0 {
		return true, scattered, m.Albedo
	}
	return false, vec.ZeroRay(), vec.ZeroColor()
}
