package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type pixel struct {
	r, g, b, a uint32
	str        string
}

func main() {
	images := getImages("./images/")
	f, err := os.Create("data.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	f.WriteString("<html><style> body {background-color:purple} .derp {width:1px; height: 1px; box-shadow: ")
	for i, img := range images {
		for j, pixel := range img {
			i = i
			j = j
			fmt.Println(pixel.str)
			_, err2 := f.WriteString(pixel.str + "\n")

			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}
	f.WriteString("}</style><body><div class='derp'></div></body></html>")
}

func getImages(dir string) [][]pixel {
	var images [][]pixel

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		img := loadImage(path)
		pixels := getPixel(img)
		images = append(images, pixels)
		return nil
	})
	return images
}

func loadImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	extension := filepath.Ext(path)
	if extension == ".jpg" || extension == ".jpeg" {
		img, err := jpeg.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		return img
	}

	img, err := png.Decode(f)

	if err != nil {
		log.Fatal(err)
	}
	return img
}

func getPixel(img image.Image) []pixel {
	bounds := img.Bounds()
	fmt.Println("X: ", bounds.Dx(), " Y: ", bounds.Dy())
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())

	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		r, g, b, a := img.At(x, y).RGBA()
		px := rgbaToPixel(img.At(x, y).RGBA())
		template := strconv.Itoa(x) + "px " + strconv.Itoa(y) + "px 1px 1px rgba(" + strconv.FormatUint(uint64(px.R), 10) + "," + strconv.FormatUint(uint64(px.G), 10) + "," + strconv.FormatUint(uint64(px.B), 10) + "," + strconv.FormatUint(uint64(px.A), 10) + ")"
		if i != bounds.Dx()*bounds.Dy()-1 {
			template = template + ","
		}
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
		pixels[i].str = template
	}

	return pixels
}
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

type Pixel struct {
	R int
	G int
	B int
	A int
}
