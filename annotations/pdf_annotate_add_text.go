/*
 * Annotate/mark up pages of a PDF file.
 * Adds a text annotation with a user specified string to a fixed location on every page.
 *
 * Run as: go run pdf_annotate_add_text.go input.pdf output.pdf text
 */

package main

import (
	"fmt"
	"os"

	unicommon "github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/common/license"
	pdfcore "github.com/unidoc/unipdf/v3/core"
	pdf "github.com/unidoc/unipdf/v3/model"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
Free trial license keys are available at: https://unidoc.io/
-----END UNIDOC LICENSE KEY-----
`

func init() {
	// Enable debug-level logging.
	// unicommon.SetLogger(unicommon.NewConsoleLogger(unicommon.LogLevelDebug))

	err := license.SetLicenseKey(licenseKey, `Company Name`)
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("go run pdf_annotate_add_text.go input.pdf output.pdf text\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]
	annotationText := os.Args[3]

	err := annotatePdfAddText(inputPath, outputPath, annotationText)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

// Annotate pdf file.
func annotatePdfAddText(inputPath string, outputPath string, annotationText string) error {
	unicommon.Log.Debug("Input PDF: %v", inputPath)

	// Read the input pdf file.
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	opt := &pdf.ReaderToWriterOpts{
		PageCallback: func(pageNum int, page *pdf.PdfPage) {
			// New text annotation.
			textAnnotation := pdf.NewPdfAnnotationText()
			textAnnotation.Contents = pdfcore.MakeString(annotationText)
			// The rect specifies the location of the markup.
			textAnnotation.Rect = pdfcore.MakeArray(pdfcore.MakeInteger(20), pdfcore.MakeInteger(100), pdfcore.MakeInteger(10+50), pdfcore.MakeInteger(100+50))

			// Add to the page annotations.
			page.AddAnnotation(textAnnotation.PdfAnnotation)
		},
	}

	pdfWriter, err := pdfReader.ToWriter(opt)
	if err != nil {
		return err
	}

	err = pdfWriter.WriteToFile(outputPath)
	if err != nil {
		return err
	}

	return nil
}
