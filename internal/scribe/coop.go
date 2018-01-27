package scribe

import "github.com/jung-kurt/gofpdf"

// NewCoopScribe creates a new scribe with the coop header and footer images.
func NewCoopScribe() *Scribe {
	pdf := gofpdf.New("P", "mm", "A4", link)

	pdf.SetMargins(margin, margin-2, margin)

	pdf.SetHeaderFunc(func() {
		pdf.ImageOptions("./internal/scribe/res/header.png", margin, 0, pageWidth, 0, true, gofpdf.ImageOptions{}, linkIdentifier, link)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetY(-margin)
		pdf.ImageOptions("./internal/scribe/res/footer.png", margin, -100, 30, 0, true, gofpdf.ImageOptions{}, linkIdentifier, link)
	})

	pdf.AddPage()

	return &Scribe{
		fpdf: pdf,
	}
}
