package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", path, err)
		return nil
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Printf("Error decoding image %s: %v\n", path, err)
		return nil
	}
	return img
}

func WriteImage(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", path, err)
		return
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		fmt.Printf("Error encoding image to file %s: %v\n", path, err)
	}
}

func Grayscale(img image.Image) image.Image {
	if img == nil {
		return nil
	}

	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func Resize(img image.Image) image.Image {
	if img == nil {
		return nil
	}
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}


