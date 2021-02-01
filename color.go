package main

import (
	"fmt"
)

func WriteColor(pixel_color Color) string {
	ir := int(255.999 * pixel_color.x)
	ig := int(255.999 * pixel_color.y)
	ib := int(255.999 * pixel_color.z)
	return fmt.Sprintf("%v %v %v\n", ir, ig, ib)
}
