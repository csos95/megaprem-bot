package main

import "log"

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Println(err)
	}

	err = CreateBot(config)
	if err != nil {
		log.Println(err)
	}

	err = run()
	if err != nil {
		log.Println(err)
	}
}
