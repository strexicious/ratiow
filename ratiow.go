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

func random_scene() (world *hittable.HittableList) {
	world = new(hittable.HittableList)

	ground_mat := materials.Lambertian{Albedo: vec.NewColor(0.5, 0.5, 0.5)}
	ground := sphere.NewSphere(vec.NewPoint3(0, -1000, 0), 1000, &ground_mat)
	world.Add(&ground)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chosen_mat := rand.Float64()
			center := vec.NewPoint3(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.Sub(vec.NewPoint3(4, 0.2, 0)).Norm() > 0.9 {
				var sphere_material hittable.Material

				if chosen_mat < 0.8 {
					// diffuse
					albedo := vec.RandomColor()
					sphere_material = &materials.Lambertian{Albedo: albedo}
					sphere := sphere.NewSphere(center, 0.2, sphere_material)
					world.Add(&sphere)
				} else if chosen_mat < 0.95 {
					// metal
					albedo := vec.RandomColorRange(0.5, 1)
					fuzz := rand.Float64() * 0.5
					sphere_material = &materials.Metal{Albedo: albedo, Fuzz: fuzz}
					sphere := sphere.NewSphere(center, 0.2, sphere_material)
					world.Add(&sphere)
				} else {
					// glass
					sphere_material = &materials.Dielectric{Ir: 1.5}
					sphere := sphere.NewSphere(center, 0.2, sphere_material)
					world.Add(&sphere)
				}
			}
		}
	}

	mat1 := materials.Dielectric{Ir: 1.5}
	sphere1 := sphere.NewSphere(vec.NewPoint3(0, 1, 0), 1.0, &mat1)

	mat2 := materials.Lambertian{Albedo: vec.NewColor(0.4, 0.2, 0.1)}
	sphere2 := sphere.NewSphere(vec.NewPoint3(-4, 1, 0), 1.0, &mat2)

	mat3 := materials.Metal{Albedo: vec.NewColor(0.7, 0.6, 0.5), Fuzz: 0.0}
	sphere3 := sphere.NewSphere(vec.NewPoint3(4, 1, 0), 1.0, &mat3)

	world.Add(&sphere1, &sphere2, &sphere3)

	return world
}

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

	const (
		aspect_ratio      = 3.0 / 2.0
		width             = 1200
		height            = int(width / aspect_ratio)
		samples_per_pixel = 500
		max_depth         = 50
	)

	// World

	world := random_scene()

	// Camera

	lookfrom := vec.NewPoint3(13, 2, 3)
	lookat := vec.NewPoint3(0, 0, 0)
	vup := vec.NewVec3(0, 1, 0)
	dist_to_focus := 10.0
	aperture := 0.1
	cam := camera.New(lookfrom, lookat, vup, 20, aspect_ratio, aperture, dist_to_focus)

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
