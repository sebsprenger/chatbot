package plugin

import (
	"strings"

	"github.com/sebsprenger/chatterschool/shared"
)

const NOTHING = ""

type Bot struct {
}

func (bot Bot) Respond(message shared.Message) string {
	if strings.HasSuffix(message.Text, "?") {
		return "Das habe ich nicht verstanden :("
	}
	return NOTHING
}
