package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

func help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	sendMessage(s, m.ChannelID, bot.help)
}

func messageLifetime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	switch len(args) {
	case 0:
		sendMessage(s, m.ChannelID, fmt.Sprintf("The current message lifetime is %d seconds.", bot.messageLifetime/time.Second))
	case 1:
		duration, err := strconv.Atoi(args[0])
		if err != nil {
			sendMessage(s, m.ChannelID, "Please use an integer for the message lifetime.")
			return
		}

		bot.messageLifetime = time.Second * time.Duration(duration)
		sendMessage(s, m.ChannelID, fmt.Sprintf("The message lifetime was set to %d seconds.", bot.messageLifetime/time.Second))
	default:
		sendMessage(s, m.ChannelID, "Too many arguments.")
	}
}

func prefix(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	switch len(args) {
	case 0:
		sendMessage(s, m.ChannelID, "Not enough arguments.")
	case 1:
		bot.config.Prefix = args[0]
		createHelp()
		sendMessage(s, m.ChannelID, fmt.Sprintf("The prefix was set to %s.", bot.config.Prefix))
	default:
		sendMessage(s, m.ChannelID, "Too many arguments.")
	}
}
