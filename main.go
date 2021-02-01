package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cheggaaa/pb"
)

func main() {
	img_width := 256
	img_height := 256
	f, err := os.OpenFile("test.ppm", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(0)
	// fmt.Printf("P3\n%v %v\n255\n", img_width, img_height)
	log.Printf("P3\n%v %v\n255\n", img_width, img_height)
	bar := pb.StartNew(img_height)
	for j := img_height - 1; j >= 0; j-- {
		bar.Increment()
		for i := 0; i < img_width; i++ {
			pixel_color := Color{float64(i) / float64(img_width-1), float64(j) / float64(img_height-1), 0.25}
			val := WriteColor(pixel_color)

			log.Print(val)
		}
	}
	bar.Finish()
	fmt.Println("Done.")

}
