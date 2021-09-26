package camera

import (
	"math"

	"github.com/strexicious/ratiow/utils"
	"github.com/strexicious/ratiow/vec"
)

type Camera struct {
	origin, lower_left_corner vec.Point3
	horizontal, vertical      vec.Vec3
	u, v, w                   vec.Vec3
	lens_radius               float64
}

func New(lookfrom, lookat vec.Point3, vup vec.Vec3, vfov, aspect_ratio, aperture, focus_dist float64) (cam *Camera) {
	theta := utils.Deg2Radians(vfov)
	h := math.Tan(theta / 2)
	viewport_height := 2 * h
	viewport_width := aspect_ratio * viewport_height

	cam = new(Camera)
	cam.w = lookfrom.Sub(lookat).Normalised()
	cam.u = vup.Cross(cam.w).Normalised()
	cam.v = cam.w.Cross(cam.u)

	cam.origin = lookfrom
	cam.horizontal = cam.u.Scale(viewport_width).Scale(focus_dist)
	cam.vertical = cam.v.Scale(viewport_height).Scale(focus_dist)
	cam.lower_left_corner = cam.origin.
		Sub(cam.horizontal.Unscale(2.0)).
		Sub(cam.vertical.Unscale(2.0)).
		Sub(cam.w.Scale(focus_dist))
	cam.lens_radius = aperture / 2
	return
}

func (cam *Camera) GetRay(s, t float64) vec.Ray {
	rd := vec.RandomInUnitDisk().Scale(cam.lens_radius)
	offset := cam.u.Scale(rd.X()).Add(cam.v.Scale(rd.Y()))

	return vec.NewRay(
		cam.origin.Add(offset),
		cam.lower_left_corner.
			Add(cam.horizontal.Scale(s)).
			Add(cam.vertical.Scale(t)).
			Sub(cam.origin).
			Sub(offset),
	)
}
