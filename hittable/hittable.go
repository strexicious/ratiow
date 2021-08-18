package hittable

import "github.com/strexicious/ratiow/vec"

type HitRecord struct {
	P         vec.Point3
	Normal    vec.Vec3
	T         float64
	FrontFace bool
}
type Hittable interface {
	Hit(r vec.Ray, t_min, t_max float64) (hit bool, rec *HitRecord)
}

type HittableList struct {
	objects []Hittable
}

func (hr *HitRecord) SetFaceNormal(r vec.Ray, outward_normal vec.Vec3) {
	hr.FrontFace = r.Direction().Dot(outward_normal) < 0
	if hr.FrontFace {
		hr.Normal = outward_normal
	} else {
		hr.Normal = outward_normal.Neg()
	}
}

func (hl *HittableList) Clear() {
	hl.objects = nil
}

func (hl *HittableList) Add(hs ...Hittable) {
	hl.objects = append(hl.objects, hs...)
}

func (hl *HittableList) Hit(r vec.Ray, t_min, t_max float64) (hit bool, rec *HitRecord) {
	hit = false
	closest_so_far := t_max

	for _, object := range hl.objects {
		if it_hit, hr := object.Hit(r, t_min, closest_so_far); it_hit {
			hit = true
			closest_so_far = hr.T
			rec = hr
		}
	}

	return
}
