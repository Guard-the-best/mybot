package main

import (
	"log"

	"github.com/Guard-the-best/mybot/internal/util"
	"github.com/Guard-the-best/mybot/internal/controller"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func main() {
	util.DefaultConfig.PrintConfig()
	bot := buildBot(&util.DefaultConfig)

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	handlerList := controller.GetHandlerList()
	bindingHandlers(bh, &handlerList)

	bh.Start()
}

func bindingHandlers(bh *th.BotHandler, handlerList *[]controller.HandlerRegister) {
	log.Println("注册handler")
	for _, handlerRegister := range *handlerList {
		log.Printf("%v: %v \n", handlerRegister.Handler, handlerRegister.Predicates)
		bh.Handle(handlerRegister.Handler, handlerRegister.Predicates...)
	}
}

func buildBot(config *util.Config) *telego.Bot {
	httpClient := fasthttp.Client{}
	if config.Network.NeedProxy() {
		if config.Network.GetProxyProtocal() == "socks5" {
			log.Println("using socks5 proxy")
			httpClient = fasthttp.Client{Dial: fasthttpproxy.FasthttpSocksDialer(config.Network.Proxy)}
		} else {
			httpClient = fasthttp.Client{Dial: fasthttpproxy.FasthttpHTTPDialer(config.Network.Proxy)}
		}
	}
	bot, err := telego.NewBot(
		config.Bot.Token,
		telego.WithFastHTTPClient(&httpClient),
	)
	if err != nil {
		log.Fatalln(0, err)
	}
	return bot
}
