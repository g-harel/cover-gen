package scribe

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"golang.org/x/text/encoding/charmap"
)

// Commonly used default values.
var align = "B"
var border = ""
var fill = false
var font = "Arial"
var fontSize = 10.0
var indent = "     "
var lineHeight = 5.0
var link = ""
var linkIdentifier = 0
var ln = 1
var margin = 18.0
var pageWidth = 210 - 2*margin
var spacer = 5.0

// Scribe provides convenient functions to write content to a pdf document.
// All methods on this type should reset to the default formatting before returning.
type Scribe struct {
	fpdf *gofpdf.Fpdf
}

// NewScribe creates a new Scribe struct.
func NewScribe() *Scribe {
	pdf := gofpdf.New("P", "mm", "A4", link)

	pdf.SetMargins(margin, margin-2, margin)

	pdf.AddPage()

	return &Scribe{
		fpdf: pdf,
	}
}

// Spacer adds a space to the document.
func (s *Scribe) Spacer(multiplier float64) {
	s.fpdf.CellFormat(pageWidth, spacer*multiplier, "", border, ln, align, fill, linkIdentifier, link)
}

// Paragraph adds a paragraph to the document.
func (s *Scribe) Paragraph(content string) {
	s.fpdf.MultiCell(pageWidth, lineHeight, f(indent+content), border, align, fill)
}

// Link adds a link to the document.
func (s *Scribe) Link(content, link string) {
	s.fpdf.SetTextColor(0, 0, 255)
	s.fpdf.SetFont(font, "U", fontSize)
	s.fpdf.CellFormat(pageWidth, lineHeight, f(content), border, ln, align, fill, linkIdentifier, link)
	s.fpdf.SetTextColor(0, 0, 0)
	s.fpdf.SetFont(font, "", fontSize)
}

// Title adds a title to the document.
func (s *Scribe) Title(content string, multiplier float64) {
	size := fontSize * multiplier
	s.fpdf.SetFont(font, "B", size)
	s.fpdf.CellFormat(pageWidth, size/2, f(content), border, ln, "CB", fill, linkIdentifier, link)
	s.fpdf.SetFont(font, "", fontSize)
}

// Line adds a line of text to the document.
func (s *Scribe) Line(content string) {
	s.fpdf.CellFormat(pageWidth, lineHeight, f(content), border, ln, align, fill, linkIdentifier, link)
}

// FormattedLine adds a line to the document with specific formatting.
func (s *Scribe) FormattedLine(content, format string) {
	s.fpdf.SetFont(font, format, fontSize)
	s.fpdf.CellFormat(pageWidth, lineHeight, f(content), border, ln, align, fill, linkIdentifier, link)
	s.fpdf.SetFont(font, "", fontSize)
}

// Save writes the document to the root folder.
func (s *Scribe) Save(name string) {
	filename := name + ".pdf"
	err := s.fpdf.OutputFileAndClose(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(filename)
}

// Encodes special characters (accents).
func f(str string) string {
	s, _ := charmap.Windows1252.NewEncoder().String(str)
	return s
}
