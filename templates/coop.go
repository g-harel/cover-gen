package main

import (
	"regexp"
	"strconv"
	"time"

	"github.com/g-harel/cover-gen/internal/scribe"
)

// ID will also be used as the output file name.
var id = "12345"

var title = "Fizz Buzz Intern"

// Contents can be copied directly from the "Company Information" section on a coop listing.
var listing = `
Organization: Lorem Ipsum Inc.
Salutation: Mr.
Job Contact First Name: John
Job Contact Last Name: Doe
Address Line One: 1234 Steve's St.
Address Line Two: 2nd floor
City: Montreal
Province / State: Québec
Postal Code / Zip Code: H1A 2B3
`

func main() {
	s := scribe.NewCoopScribe()

	date := getFormattedDate()
	organization, name, greeting, address, location := getListingData(listing)

	s.Title("Your Name Here", 2)
	s.Title("Portfolio/Website link", 1)

	s.Spacer(2)
	s.Line(date)

	s.Spacer(1)
	s.Line(name)
	s.Line(organization)
	s.Line(address)
	s.Line(location)

	s.Spacer(2)
	s.FormattedLine(title, "B")

	s.Spacer(1)
	s.Line(greeting)

	s.Spacer(2)
	s.Paragraph("Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam.")

	s.Spacer(1)
	s.Paragraph("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.")

	s.Spacer(1)
	s.Link("Your Website Demo", "https://example.org/project")

	s.Spacer(1)
	s.Paragraph("At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus.")

	s.Spacer(1)
	s.Paragraph("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")

	s.Spacer(2)
	s.Line("Your Name Here")

	s.Save(id)
}

func getFormattedDate() string {
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

	return month.String() + " " + strconv.Itoa(day) + dayPostfix + " " + strconv.Itoa(year)
}

// Parses and formats listing information.
func getListingData(listing string) (organization, name, greeting, address, location string) {
	coopRegExp := regexp.MustCompile("(?s)" + ".*" +
		"Organization: ([^\n]+).*" +
		"Salutation: ([^\n]+).*" +
		"Job Contact First Name: ([^\n]+).*" +
		"Job Contact Last Name: ([^\n]+).*" +
		"Address Line One: ([^\n]+).*" +
		"City: ([^\n]+).*" +
		"Province / State: ([^\n]+).*" +
		"Postal Code / Zip Code: ([^\n]+).*")

	matchResult := coopRegExp.FindStringSubmatch(listing)

	if matchResult == nil {
		panic("could not parse listing")
	}

	organization = matchResult[1]
	name = matchResult[3] + " " + matchResult[4]
	greeting = "Dear " + matchResult[2] + " " + matchResult[4] + ","
	address = matchResult[5]
	location = matchResult[6] + ", " + matchResult[7] + " " + matchResult[8]

	return
}
