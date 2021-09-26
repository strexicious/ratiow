package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/strexicious/ratiow/camera"
	"github.com/strexicious/ratiow/hittable"
	"github.com/strexicious/ratiow/hittable/materials"
	"github.com/strexicious/ratiow/sphere"
	"github.com/strexicious/ratiow/vec"
)

func ray_color(r vec.Ray, world *hittable.HittableList, depth int32) vec.Color {
	// If we've exceeded the ray bounce limit, no more light is gathered.
	if depth <= 0 {
		return vec.NewColor(0, 0, 0)
	}

	if hit, hr := world.Hit(r, 0.001, math.Inf(1)); hit {
		did_scatter, scattered, attenuation := hr.Mat.Scatter(r, hr)
		if did_scatter {
			return ray_color(scattered, world, depth-1).ComponentWiseScale(attenuation)
		}

		return vec.ZeroColor()
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
	const max_depth = 50

	// World

	ground_mat := materials.Lambertian{Albedo: vec.NewColor(0.8, 0.8, 0.0)}
	center_mat := materials.Lambertian{Albedo: vec.NewColor(0.1, 0.2, 0.5)}
	left_mat := materials.Dielectric{Ir: 1.5}
	right_mat := materials.Metal{Albedo: vec.NewColor(0.8, 0.6, 0.2), Fuzz: 0.0}

	spheres := []sphere.Sphere{
		sphere.NewSphere(vec.NewPoint3(0, -100.5, -1), 100, &ground_mat),
		sphere.NewSphere(vec.NewPoint3(0, 0, -1), 0.5, &center_mat),
		sphere.NewSphere(vec.NewPoint3(-1, 0, -1), 0.5, &left_mat),
		sphere.NewSphere(vec.NewPoint3(-1, 0, -1), -0.45, &left_mat),
		sphere.NewSphere(vec.NewPoint3(1, 0, -1), 0.5, &right_mat),
	}

	world := new(hittable.HittableList)
	world.Add(&spheres[0], &spheres[1], &spheres[2], &spheres[3], &spheres[4])

	// Camera

	cam := camera.New(vec.NewPoint3(-2, 2, 1), vec.NewPoint3(0, 0, -1), vec.NewVec3(0, 1, 0), 20, aspect_ratio)

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
				pixel_color = pixel_color.Add(ray_color(r, world, max_depth))
			}
			pixel_color.WriteColor(os.Stdout, samples_per_pixel)
		}
	}

	println("\nDone.")
}
