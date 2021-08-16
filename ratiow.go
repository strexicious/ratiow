package main

import (
	"fmt"
	"os"

	"github.com/strexicious/ratiow/vec"
)

func ray_color(r vec.Ray) vec.Color {
	unit_direction := r.Direction().Normalised()
	t := 0.5 * (unit_direction.Y() + 1.0)
	return vec.NewColor(1.0, 1.0, 1.0).Scale(1.0 - t).Add(vec.NewColor(0.5, 0.7, 1.0).Scale(t))
}

func main() {

	// Image

	const aspect_ratio = 16.0 / 9.0
	const width = 400
	const height = int(width / aspect_ratio)

	// Camera

	const viewport_height = 2.0
	const viewport_width = aspect_ratio * viewport_height
	const focal_length = 1.0

	origin := vec.NewPoint3(0, 0, 0)
	horizontal := vec.NewVec3(viewport_width, 0, 0)
	vertical := vec.NewVec3(0, viewport_height, 0)
	lower_left_corner := origin.
		Sub(horizontal.Unscale(2.0)).
		Sub(vertical.Unscale(2.0)).
		Sub(vec.NewVec3(0, 0, focal_length))

	// Render

	fmt.Printf("P3\n%d %d\n255\n", width, height)

	for j := height - 1; j >= 0; j-- {
		print("\rScanlines remaining: ", j, " ")
		for i := 0; i < width; i++ {
			u := float64(i) / float64(width-1)
			v := float64(j) / float64(height-1)
			r := vec.NewRay(origin, lower_left_corner.Add(horizontal.Scale(u)).Add(vertical.Scale(v)).Sub(origin))
			c := ray_color(r)
			c.WriteColor(os.Stdout)
		}
	}

	println("\nDone.")
}
