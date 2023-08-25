package main

import (
	"fmt"
	"hello_bot/bot"
	"hello_bot/config"
)

func main() {
	fmt.Println("Bot is running")

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	//entry chanel / notifiation chanel
	<-make(chan struct{})

}
