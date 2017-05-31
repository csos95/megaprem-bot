package megaprem_bot

import "log"

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Println(err)
	}

	bot, err := NewBot(config)
	if err != nil {
		log.Println(err)
	}

	err = bot.run()
	if err != nil {
		log.Println(err)
	}
}
