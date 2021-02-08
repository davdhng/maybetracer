package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/cheggaaa/pb"
)

func rayColor(r Ray, world HittableList, depth int) Vec3 {
	// var rec HitRecord
	if depth <= 0 {
		return Vec3{0, 0, 0}
	}
	if hit, rec := world.Hit(r, 0.001, math.MaxFloat64); hit {
		// target := rec.p.Add(rec.normal).Add(rec.p.randSphere())
		target := rec.p.Add(rec.normal).Add(RandUnitVec())
		return (rayColor(Ray{rec.p, target.Sub(rec.p)}, world, depth-1)).Scale(0.5)
		// return (rec.normal.Add(Vec3{1.0, 1.0, 1.0})).Scale(0.5)
	}
	unitDirection := r.Direction.unitVector()
	t := 0.5 * (unitDirection.y + 1.0)
	return (Vec3{1.0, 1.0, 1.0}.Scale(1.0 - t)).Add(Vec3{0.5, 0.7, 1.0}.Scale(t))
}

func main() {
	imgWidth := 400
	aspectRatio := 16.0 / 9.0
	imgHeight := int(float64(imgWidth) / aspectRatio)
	samplesPerPixel := 100
	maxDepth := 50

	var world HittableList
	world.Add(Sphere{Vec3{0, 0, -1}, 0.5})
	world.Add(Sphere{Vec3{0, -100.5, -1}, 100})

	cam := NewCamera()

	f, err := os.OpenFile("test.ppm", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(0)
	log.Printf("P3\n%v %v\n255\n", imgWidth, imgHeight)
	bar := pb.StartNew(imgHeight)

	for j := imgHeight; j > 0; j-- {
		for i := 0; i < imgWidth; i++ {
			pixelColor := Vec3{0, 0, 0}
			rnd := rand.New(rand.NewSource(int64(42 * j)))
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rnd.Float64()) / float64(imgWidth-1)
				v := (float64(j) + rnd.Float64()) / float64(imgHeight-1)
				r := cam.getRay(u, v)
				pc := rayColor(r, world, maxDepth)
				pixelColor = pixelColor.Add(pc)
			}
			val := WriteColor(pixelColor, samplesPerPixel)

			log.Print(val)
		}
		bar.Increment()
	}
	bar.Finish()
	fmt.Println("Done.")

}
