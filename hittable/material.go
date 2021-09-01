package hittable

import "github.com/strexicious/ratiow/vec"

type Material interface {
	Scatter(r_in vec.Ray, rec *HitRecord) (bool, vec.Ray, vec.Color)
}
