package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func sendMessage(s *discordgo.Session, channelID string, message string) {
	msg, err := s.ChannelMessageSend(channelID, fmt.Sprintf("```\n%s\n```", message))
	if err != nil {
		log.Println(err)
	}

	if bot.messageLifetime != 0 {
		lifetimeChan := time.After(bot.messageLifetime)
		go func() {
			<-lifetimeChan
			deleteMessage(s, channelID, msg)
		}()
	}
}

func deleteMessage(s *discordgo.Session, channelID string, message *discordgo.Message) {
	s.ChannelMessageDelete(channelID, message.ID)
}
