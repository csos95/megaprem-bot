package main

import (
	"fmt"
)

func (b *Bot) sendMessage(channelID string, message string) {
	bot.ChannelMessageSend(channelID, fmt.Sprintf("```\n%s\n```", message))
}
