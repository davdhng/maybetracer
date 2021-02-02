package main

import (
	"fmt"
	"math"
)

func clamp(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

func mapp(value, inFrom, inTo, outFrom, outTo float64) float64 {
	// restrict value to input interval
	value = clamp(value, inFrom, inTo)
	return outTo*(value-inFrom)/(inTo-inFrom) + outFrom
}

func WriteColor(pixel_color Vec3, samples_per_pixel int) string {
	r := pixel_color.x
	g := pixel_color.y
	b := pixel_color.z

	scale := 1.0 / float64(samples_per_pixel)
	r = math.Sqrt(scale * r)
	g = math.Sqrt(scale * g)
	b = math.Sqrt(scale * b)

	maxColor := 255.0

	ir := uint8(mapp(r, 0, 1, 0, maxColor))
	ig := uint8(mapp(g, 0, 1, 0, maxColor))
	ib := uint8(mapp(b, 0, 1, 0, maxColor))

	// ir := int(256 * clamp(r, 0.0, 0.999))
	// ig := int(256 * clamp(g, 0.0, 0.999))
	// ib := int(256 * clamp(b, 0.0, 0.999))

	return fmt.Sprintf("%v %v %v\n", ir, ig, ib)
}
