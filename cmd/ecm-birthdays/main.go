package main

import (
	"github.com/pa3ng/ecm-birthdays/internal/gmail"
	// "github.com/pa3ng/ecm-birthdays/internal/sheets"
)

func main() {
	// sheets.PrintResidentNamesAndGenders()
	// gmail.PrintMyLabels()
	gmail.SendEmail("raphael.santodomingo@gpmail.org")
}
