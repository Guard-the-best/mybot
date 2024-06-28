package util

import (
	"log"

	th "github.com/mymmrac/telego/telegohandler"
)

type HandlerRegister struct {
	Handler    th.Handler
	Predicates []th.Predicate
}

var handlerList = []HandlerRegister{}

// 防止注册方修改hadler以及predicate，此处不用指针
func RegisterHandler(handler th.Handler, predicates ...th.Predicate) {
	log.Printf("注册 %v", handler)
	handlerRegister := HandlerRegister{Handler: handler, Predicates: predicates}
	handlerList = append(handlerList, handlerRegister)
}

// 防止外部修改，返回一个复制的handlerList
func GetHandlerList() []HandlerRegister {
	return handlerList
}

func BindingHandlers(bh *th.BotHandler) {
	log.Println("注册handler")
	for _, handlerRegister := range handlerList {
		log.Printf("%v: %v \n", handlerRegister.Handler, handlerRegister.Predicates)
		bh.Handle(handlerRegister.Handler, handlerRegister.Predicates...)
	}
}