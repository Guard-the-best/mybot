package util

import (
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var HelpDoc = ""

func AddHelpDoc(command string, comment string) {
	// log.Println(command + ":" + comment)
	if HelpDoc != "" {
		HelpDoc += "\n"
	}
	if command != "" {
		HelpDoc += fmt.Sprintf("`/%s` :  %s", command, comment)
	} else {
		HelpDoc += comment
	}
	log.Print(HelpDoc)
}

func init() {
	AddHelpDoc("help", "获得一些帮助")
}

func GetHelp(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		HelpDoc,
	).WithParseMode(telego.ModeMarkdownV2))
}
