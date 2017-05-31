package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (b *Bot) sendMessage(s *discordgo.Session, channelID string, message string) {
	log.Println(message)
	s.ChannelMessageSend(channelID, fmt.Sprintf("```\n%s\n```", message))
}
