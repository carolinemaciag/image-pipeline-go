package imageprocessing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestReadImage(t *testing.T) {
	// Read image from the test folder
	img := ReadImage("../test/image1.jpg")
	// Assuming ReadImage doesn't return an error, so we just check if img is not nil
	assert.NotNil(t, img, "Expected an image, but got nil")
}

func TestWriteImage(t *testing.T) {
	// Read image from the test folder
	img := ReadImage("../test/image1.jpg")
	if img == nil {
		t.Errorf("Failed to read test image")
		return
	}

	// Test WriteImage
	outputPath := "../test/output/test_output.jpg"
	err := WriteImage(outputPath, img)
	if err != nil {
		t.Errorf("Failed to write test image: %v", err)
	}

	// Check if output file was created
	_, err = os.Stat(outputPath)
	if os.IsNotExist(err) {
		t.Errorf("Output file was not created")
	}
}

func TestGrayscale(t *testing.T) {
	// Read image from the test folder
	img := ReadImage("../test/image1.jpg")
	if img == nil {
		t.Errorf("Failed to read test image")
		return
	}

	// Apply grayscale filter
	grayImg := Grayscale(img)
	// Check if the image is grayscale (basic check)
	assert.NotNil(t, grayImg, "Expected a grayscale image, but got nil")
}

func TestResize(t *testing.T) {
	// Read image from the test folder
	img := ReadImage("../test/image1.jpg")
	if img == nil {
		t.Errorf("Failed to read test image")
		return
	}

	// Resize image
	resizedImg := Resize(img)
	// Check if the resized image is valid
	assert.NotNil(t, resizedImg, "Expected a resized image, but got nil")
}











