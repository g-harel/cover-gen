package main

import (
	"strconv"
	"time"

	"golang.org/x/text/encoding/charmap"

	"github.com/jung-kurt/gofpdf"
)

var border = ""
var fontSize = 10.0
var lineHeight = 5.0
var margin = 18.0
var pageWidth = 210 - 2*margin
var spacer = 6.0
var indent = "     "

func f(str string) string {
	s, _ := charmap.Windows1252.NewEncoder().String(str)
	return s
}

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetMargins(margin, margin-2, margin)

	pdf.SetHeaderFunc(func() {
		pdf.ImageOptions("./res/header.png", margin, 0, pageWidth, 0, true, gofpdf.ImageOptions{}, 0, "")
	})

	pdf.SetFooterFunc(func() {
		pdf.SetY(-margin)
		pdf.ImageOptions("./res/footer.png", margin, -100, 30, 0, true, gofpdf.ImageOptions{}, 0, "")
	})

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(pageWidth, 12, f("Gabriel Harel"), border, 1, "CB", false, 0, "")

	pdf.SetFont("Arial", "B", fontSize)
	pdf.CellFormat(pageWidth, 6, f("github.com/g-harel"), border, 1, "CB", false, 0, "")

	pdf.SetFont("Arial", "", fontSize)
	year, month, day := time.Now().Date()
	var dayPostfix string
	switch day {
	case 1, 21, 31:
		dayPostfix = "st"
	case 2, 22:
		dayPostfix = "nd"
	case 3, 23:
		dayPostfix = "rd"
	default:
		dayPostfix = "th"
	}
	pdf.CellFormat(pageWidth, 2*spacer, "", border, 1, "B", false, 0, "")
	pdf.CellFormat(pageWidth, lineHeight, f(month.String()+" "+strconv.Itoa(day)+dayPostfix+" "+strconv.Itoa(year)), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.CellFormat(pageWidth, lineHeight, f("Nicholas Osadchuck"), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, lineHeight, f("Vigilant Global"), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, lineHeight, f("800 René-Lévesque Blvd"), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, lineHeight, f("Montréal, Québec H3B 1X9"), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, 2*spacer, "", border, 1, "B", false, 0, "")
	pdf.SetFont("Arial", "B", fontSize)
	pdf.CellFormat(pageWidth, lineHeight, f("Front-End Web Developer - Internship"), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.SetFont("Arial", "", fontSize)
	pdf.CellFormat(pageWidth, lineHeight, f("Dear Mr. Osadchuck,"), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, 2*spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f(indent+"I am a Software Engineering student at Concordia in the Institute for Co-operative Education that is very interested in web development. Ever since I first started learning about the web, I have been hooked by the fast-paced ecosystem and the large community. This interest has driven me to create many web-based projects that haven given me a good foundation in JavaScript, HTML and CSS, but have also allowed me to explore React, NodeJS, Gulp and other modern technologies. I encourage you to visit my GitHub repositories where I have links to live demos of these projects."), border, "B", false)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.SetTextColor(0, 0, 255)
	pdf.SetFont("Arial", "U", fontSize)
	pdf.CellFormat(pageWidth, lineHeight, f("https://github.com/g-harel"), border, 1, "B", false, 0, "https://github.com/g-harel")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", fontSize)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f(indent+"My previous work term in Ubisoft’s CRM team has given me valuable work experience with working in an international team and dealing with strict deadlines. During this internship, I also built tools to facilitate work within the team. This taught me how to implement a client’s requirements and add features according to feedback."), border, "B", false)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f(indent+"I am very motivated to learn and it is definitely something I want to accomplish during this work term. I believe that, as an employee, I would be able to use and build upon my technical knowledge and am confident that I will be able to adapt to new concepts in a timely manner."), border, "B", false)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f("I appreciate your consideration for this position and sincerely look forward to hearing back from you."), border, "B", false)

	pdf.CellFormat(pageWidth, 2*spacer, "", border, 1, "B", false, 0, "")
	pdf.CellFormat(pageWidth, lineHeight, f("Gabriel Harel"), border, 1, "B", false, 0, "")

	pdf.OutputFileAndClose("hello.pdf")
}
