# Image Processing Pipeline with and without Goroutines

This Go project is based off of Amrit Singh's go_21_goroutines_pipeline repository- which implements an image processing pipeline that loads images, converts them to grayscale, resizes them, and saves the processed images. The project can be run with or without goroutines to compare processing times for concurrent versus sequential execution.

## Features

- **Load Image**: Reads images from the specified file paths.
- **Convert to Grayscale**: Converts each image to grayscale.
- **Resize**: Resizes the image to a predefined size.
- **Save Image**: Saves the processed image to an output location.

The program can run both with and without the use of goroutines for concurrent execution. You can compare the processing times.

### Benchmarking
The program tracks the time spent on each stage of the image processing pipeline:
- **Loading images**
- **Converting images to grayscale**
- **Resizing images**
- **Saving images**

At the end of each run, the total processing time for the entire pipeline is displayed. This allows you to compare the performance with and without goroutines.

## How to Run the Program

### Running with Goroutines

To run the image processing pipeline using goroutines for parallel processing:

```bash
go run main.go -goroutines=true
```

### Running without Goroutines

To run the image processing pipeline without using goroutines:

```bash
go run main.go -goroutines=false
```

## Unit testing was attempted, however this portion ended in failure.
- ChatGPT was used in attempting to troubleshoot errors in this section, however after multiple attempts, the same errors persisted.
