package camera

import "github.com/strexicious/ratiow/vec"

type Camera struct {
	origin, lower_left_corner vec.Point3
	horizontal, vertical      vec.Vec3
}

func DefaultCamera() (cam *Camera) {
	const aspect_ratio = 16.0 / 9.0
	const viewport_height = 2.0
	const viewport_width = aspect_ratio * viewport_height
	const focal_length = 1.0

	cam = new(Camera)
	cam.horizontal = vec.NewVec3(viewport_width, 0, 0)
	cam.vertical = vec.NewVec3(0, viewport_height, 0)
	cam.origin = vec.NewPoint3(0, 0, 0)
	cam.lower_left_corner = cam.origin.
		Sub(cam.horizontal.Unscale(2.0)).
		Sub(cam.vertical.Unscale(2.0)).
		Sub(vec.NewVec3(0, 0, focal_length))

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
