package main

import (
	"flag"
	"log"
	"os"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	url := flag.String("url", "", "URL of article to convert")
	file := flag.String("o", "doc.pdf", "Output pdf filename")
	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(3)
	}

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage(*url)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(95.50)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(*file)
	if err != nil {
		log.Fatal(err)
	}
}
