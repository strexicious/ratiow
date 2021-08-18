package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/strexicious/ratiow/camera"
	"github.com/strexicious/ratiow/hittable"
	"github.com/strexicious/ratiow/sphere"
	"github.com/strexicious/ratiow/vec"
)

func ray_color(r vec.Ray, world *hittable.HittableList) vec.Color {
	if hit, hr := world.Hit(r, 0, math.Inf(1)); hit {
		return vec.NewColor(1, 1, 1).Add(hr.Normal).Scale(0.5)
	}

	unit_direction := r.Direction().Normalised()
	t := 0.5 * (unit_direction.Y() + 1.0)
	return vec.NewColor(1.0, 1.0, 1.0).Scale(1.0 - t).Add(vec.NewColor(0.5, 0.7, 1.0).Scale(t))
}

func main() {

	// Image

	const aspect_ratio = 16.0 / 9.0
	const width = 400
	const height = int(width / aspect_ratio)
	const samples_per_pixel = 100

	// World

	spheres := []sphere.Sphere{
		sphere.NewSphere(vec.NewPoint3(0, 0, -1), 0.5),
		sphere.NewSphere(vec.NewPoint3(0, -100.5, -1), 100),
	}

	world := new(hittable.HittableList)
	world.Add(&spheres[0], &spheres[1])

	// Camera

	cam := camera.DefaultCamera()

	// Render

	fmt.Printf("P3\n%d %d\n255\n", width, height)

	for j := height - 1; j >= 0; j-- {
		print("\rScanlines remaining: ", j, " ")
		for i := 0; i < width; i++ {
			pixel_color := vec.ZeroColor()
			for s := 0; s < samples_per_pixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(width-1)
				v := (float64(j) + rand.Float64()) / float64(height-1)
				r := cam.GetRay(u, v)
				pixel_color = pixel_color.Add(ray_color(r, world))
			}
			pixel_color.WriteColor(os.Stdout, samples_per_pixel)
		}
	}

	println("\nDone.")
}
