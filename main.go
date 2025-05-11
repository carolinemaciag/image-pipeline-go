package main

import (
	"fmt"
	"image"
	"flag"
	"strings"
	"time"
	imageprocessing "goroutines_pipeline/image_processing"
)

// Job represents the unit of work
type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

// Sequential pipeline stages (without goroutines)
func loadImageSequential(paths []string) []Job {
	var jobs []Job
	for _, p := range paths {
		job := Job{
			InputPath: p,
			OutPath:   strings.Replace(p, "images/", "images/output/", 1),
		}
		job.Image = imageprocessing.ReadImage(p)
		jobs = append(jobs, job)
	}
	return jobs
}

func convertToGrayscaleSequential(jobs []Job) []Job {
	for i := range jobs {
		jobs[i].Image = imageprocessing.Grayscale(jobs[i].Image)
	}
	return jobs
}

func resizeSequential(jobs []Job) []Job {
	for i := range jobs {
		jobs[i].Image = imageprocessing.Resize(jobs[i].Image)
	}
	return jobs
}

func saveImageSequential(jobs []Job) {
	for _, job := range jobs {
		imageprocessing.WriteImage(job.OutPath, job.Image)
	}
}

// Pipeline with goroutines
func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, p := range paths {
			job := Job{
				InputPath: p,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
			}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {
	// Flag to control if goroutines should be used
	useGoroutines := flag.Bool("goroutines", true, "Run the pipeline with goroutines (default: true)")
	flag.Parse()

	// Image paths
	imagePaths := []string{
		"images/image1.jpg",
		"images/image2.jpg",
		"images/image3.jpg",
		"images/image4.jpg",
	}

	// Measure the execution time with goroutines or sequential
	var startTime time.Time
	var endTime time.Time

	if *useGoroutines {
		// Run with goroutines
		startTime = time.Now()
		channel1 := loadImage(imagePaths)
		channel2 := convertToGrayscale(channel1)
		channel3 := resize(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if success {
				fmt.Println("Image processed successfully!")
			} else {
				fmt.Println("Image processing failed.")
			}
		}
		endTime = time.Now()
	} else {
		// Run sequentially (no goroutines)
		startTime = time.Now()
		jobs := loadImageSequential(imagePaths)
		jobs = convertToGrayscaleSequential(jobs)
		jobs = resizeSequential(jobs)
		saveImageSequential(jobs)
		endTime = time.Now()
	}

	// Output the time taken for the pipeline
	fmt.Printf("Total pipeline processing took: %v\n", endTime.Sub(startTime))
}


