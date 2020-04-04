/*
 * This example showcases the decompression of the jbig2 encoded image and storing into
 * commonly used jpg format.
 */

package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidoc/unipdf/v3/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run jbig2_decompress_image.go img.jb2 ...\n")
		os.Exit(1)
	}
	inputImage := os.Args[1]
	_, fileName := filepath.Split(inputImage)

	f, err := os.Open(inputImage)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Create new JBIG2 Encoder/Decoder context.
	enc := &core.JBIG2Encoder{}

	// Decode all images from the 'data'.
	images, err := enc.DecodeImages(data)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	// The 'checkerboard-squares-black-white.jb2 file should have exactly one image stored.
	if len(images) != 1 {
		log.Fatalf("Error: Only a single image should be decoded\n")
	}
	fileNameWithoutExtension := func(filename string) string {
		if i := strings.LastIndex(filename, "."); i != -1 {
			return filename[:i]
		}
		return filename
	}
	// Create a new file for the decoded image.
	dec, err := os.Create(fileNameWithoutExtension(fileName) + ".jpg")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer dec.Close()
	if err = jpeg.Encode(dec, images[0], &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
