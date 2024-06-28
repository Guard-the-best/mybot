package controller

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/Guard-the-best/mybot/internal/util"
)

func helpHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), util.HelpDoc))
}

func init() {
	util.RegisterHandler(helpHandler, th.CommandEqual("help"))
}
