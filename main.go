package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/cheggaaa/pb"
)

func rayColor(r Ray) Vec3 {
	t := hitSphere(Vec3{0, 0, -1}, 0.5, r)
	if t > 0.0 {
		N := unit_vector(r.At(t).Sub(Vec3{0, 0, -1}))
		return Vec3{N.x + 1, N.y + 1, N.z + 1}.Scale(0.5)
	}
	unitDirection := unit_vector(r.Direction)
	t = 0.5 * (unitDirection.y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.Scale(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.Scale(t))
}

func hitSphere(center Vec3, radius float64, r Ray) float64 {
	oc := r.Origin.Sub(center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction) * 2.0
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}
}

func main() {
	imgWidth := 400
	aspectRatio := 16.0 / 9.0
	imgHeight := int(float64(imgWidth) / aspectRatio)

	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := Vec3{0, 0, 0}
	horizontal := Vec3{viewportWidth, 0, 0}
	vertical := Vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin.Sub(horizontal.Scale(0.5)).Sub(vertical.Scale(0.2)).Sub(Vec3{0, 0, focalLength})

	f, err := os.OpenFile("test.ppm", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(0)
	log.Printf("P3\n%v %v\n255\n", imgWidth, imgHeight)
	bar := pb.StartNew(imgHeight)

	for j := imgHeight - 1; j >= 0; j-- {
		bar.Increment()
		for i := 0; i < imgWidth; i++ {
			u := float64(i) / float64(imgWidth-1)
			v := float64(j) / float64(imgWidth-1)
			r := Ray{origin, lowerLeftCorner.Add(horizontal.Scale(u)).Add(vertical.Scale(v)).Sub(origin)}
			// pixel_color := Color{float64(i) / float64(imgWidth-1), float64(j) / float64(imgHeight-1), 0.25}
			var pixelColor Vec3 = rayColor(r)
			val := WriteColor(pixelColor)

			log.Print(val)
		}
	}
	bar.Finish()
	fmt.Println("Done.")

}
