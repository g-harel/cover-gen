package main

import (
	"fmt"
	"regexp"
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
var spacer = 5.0
var indent = "     "

func f(str string) string {
	s, _ := charmap.Windows1252.NewEncoder().String(str)
	return s
}

func match(regEx, url string) map[string]string {
	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	var found bool
	paramsMap := make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			found = true
			paramsMap[name] = match[i]
		}
	}
	if found {
		paramsMap["found"] = "true"
	}
	return paramsMap
}

var id = ""
var title = "Software Developer Intern (Front End)"
var listring = `
Organization: SoftSim Technologies Inc
Salutation: Ms.
Job Contact First Name: Farah
Job Contact Last Name: Zerrouki
Address Line One: 500 Place d'Armes
Address Line Two: suite 1800
City: Montreal
Province / State: Quebec
Postal Code / Zip Code: H2Y 2W2
`

func main() {
	listing := match("(.|\n)*"+
		"Organization: (?P<org>.+)(.|\n)*"+
		"Salutation: (?P<pronoun>.+)(.|\n)*"+
		"Job Contact First Name: (?P<firstname>.+)(.|\n)*"+
		"Job Contact Last Name: (?P<lastname>.+)(.|\n)*"+
		"Address Line One: (?P<addr>.+)(.|\n)*"+
		"City: (?P<city>.+)(.|\n)*"+
		"Province / State: (?P<province>.+)(.|\n)*"+
		"Postal Code / Zip Code: (?P<zip>.+)(.|\n)*"+
		"(.|\n)*", listring)
	if listing["found"] != "true" {
		fmt.Println("could not parse listing")
		return
	}

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
	pdf.CellFormat(pageWidth, lineHeight, f(listing["firstname"]+" "+listing["lastname"]), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, lineHeight, f(listing["org"]), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, lineHeight, f(listing["addr"]), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, lineHeight, f(listing["city"]+", "+listing["province"]+" "+listing["zip"]), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, 2*spacer, "", border, 1, "B", false, 0, "")
	pdf.SetFont("Arial", "B", fontSize)
	pdf.CellFormat(pageWidth, lineHeight, f(title), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.SetFont("Arial", "", fontSize)
	pdf.CellFormat(pageWidth, lineHeight, f("Dear "+listing["pronoun"]+" "+listing["lastname"]+","), border, 1, "B", false, 0, "")

	pdf.CellFormat(pageWidth, 2*spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f(indent+"I am a Software Engineering student at Concordia in the Institute for Co-operative Education that is very interested in web development. Ever since I first started learning about the web, I have been hooked by the fast-paced ecosystem and the large community. This interest has driven me to create many web-based projects that have given me a good foundation in JavaScript, Go, HTML, CSS and other modern tools. I encourage you to visit my GitHub repositories where these projects are hosted."), border, "B", false)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.SetTextColor(0, 0, 255)
	pdf.SetFont("Arial", "U", fontSize)
	pdf.CellFormat(pageWidth, lineHeight, f("https://github.com/g-harel"), border, 1, "B", false, 0, "https://github.com/g-harel")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", fontSize)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f(indent+"My first work term in Ubisoft’s CRM team has given me valuable work experience with working in an international team and dealing with strict deadlines. During this internship, I also built tools to facilitate work within the team. This taught me how to implement a client’s requirements and add features according to feedback."), border, "B", false)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f(indent+"My second work term as an AppDirect Frontend Platform Intern introduced me to the scale and complexity of modern software development. Working in a team alongside experienced developers was a great opportunity for me to sharpen my skills with valuable feedback and demonstrate my autonomy."), border, "B", false)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f(indent+"I am very motivated to learn and it is definitely something I want to accomplish during this work term. I believe that, as an employee, I would be able to use and build upon my technical knowledge and am confident that I will be able to adapt to new concepts in a timely manner."), border, "B", false)

	pdf.CellFormat(pageWidth, spacer, "", border, 1, "B", false, 0, "")
	pdf.MultiCell(pageWidth, lineHeight, f("I appreciate your consideration for this position and sincerely look forward to hearing back from you."), border, "B", false)

	pdf.CellFormat(pageWidth, 2*spacer, "", border, 1, "B", false, 0, "")
	pdf.CellFormat(pageWidth, lineHeight, f("Gabriel Harel"), border, 1, "B", false, 0, "")

	pdf.OutputFileAndClose(id + ".pdf")
}
