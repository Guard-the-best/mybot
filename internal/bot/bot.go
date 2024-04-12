package bot

import (
	"fmt"
	"log"

	"github.com/Guard-the-best/mybot/internal/utils"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var HelpDoc = ""

func AddHelpDoc(command string, comment string) {
	if HelpDoc != "" {
		HelpDoc += "\n"
	}
	if command != "" {
		HelpDoc += fmt.Sprintf("\\`%s` :  %s", command, comment)
	} else {
		HelpDoc += comment
	}
}

func init() {
	AddHelpDoc("help", "获得一些帮助")
	utils.RegisterHandler(
		GetHelp,
		th.CommandEqual("help"),
	)
}

func GetHelp(bot *telego.Bot, update telego.Update) {
	log.Println("我觉得我还可以再抢救一下")
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		"我觉得我还可以再抢救一下",
	).WithParseMode(telego.ModeMarkdownV2))
}
