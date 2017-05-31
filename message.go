package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) sendMessage(s *discordgo.Session, channelID string, message string) {
	s.ChannelMessageSend(channelID, fmt.Sprintf("```\n%s\n```", message))
}
