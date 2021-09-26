package camera

import (
	"math"

	"github.com/strexicious/ratiow/utils"
	"github.com/strexicious/ratiow/vec"
)

type Camera struct {
	origin, lower_left_corner vec.Point3
	horizontal, vertical      vec.Vec3
}

func New(lookfrom, lookat vec.Point3, vup vec.Vec3, vfov, aspect_ratio float64) (cam *Camera) {
	theta := utils.Deg2Radians(vfov)
	h := math.Tan(theta / 2)
	viewport_height := 2 * h
	viewport_width := aspect_ratio * viewport_height

	w := lookfrom.Sub(lookat).Normalised()
	u := vup.Cross(w).Normalised()
	v := w.Cross(u)

	cam = new(Camera)
	cam.origin = lookfrom
	cam.horizontal = u.Scale(viewport_width)
	cam.vertical = v.Scale(viewport_height)
	cam.lower_left_corner = cam.origin.
		Sub(cam.horizontal.Unscale(2.0)).
		Sub(cam.vertical.Unscale(2.0)).
		Sub(w)
	return
}

func (cam *Camera) GetRay(u, v float64) vec.Ray {
	return vec.NewRay(
		cam.origin,
		cam.lower_left_corner.
			Add(cam.horizontal.Scale(u)).
			Add(cam.vertical.Scale(v)).
			Sub(cam.origin),
	)
}
