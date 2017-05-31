package main

import "log"

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Println(err)
	}

	bot, err = CreateBot(config)
	if err != nil {
		log.Println(err)
	}

	err = bot.run()
	if err != nil {
		log.Println(err)
	}
}
