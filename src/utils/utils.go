package utils

import (
	"fmt"
	"strings"

	"github.com/paulrademacher/climenu"
)

func TrueOrFalse(requestedQuestion string, defaultQuestion bool) bool {
	var yes string
	var no string

	if defaultQuestion {
		yes = "Y"
		no = "n"
	} else {
		no = "N"
		yes = "y"
	}

	qst := fmt.Sprintf(" (%s/%s)", yes, no)

	newQuestion := requestedQuestion + qst
	question := strings.ToLower(climenu.GetText(newQuestion, "nothing"))

	if question != "y" && question != "n" {
		return defaultQuestion
	}

	return question == "y"
}
