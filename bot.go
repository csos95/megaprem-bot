package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var bot *Bot

type Bot struct {
	*discordgo.Session
}

func CreateBot(config *Config) (*Bot, error) {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	bot := &Bot{dg}

	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
	bot.AddHandler(guildCreate)

	return bot, nil
}

func (b *Bot) run() error {
	err := b.Open()
	if err != nil {
		return err
	}

	fmt.Println("Megaprem Bot is now running. Press CRTL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	b.Close()
	return nil
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "m!help")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == bot.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "r!") {
		parts := strings.Fields(m.Content)
		switch parts[0] {
		case "m!help":
			bot.sendMessage(s, m.ChannelID, helpText())
		}
	}
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			bot.sendMessage(s, channel.ID, "Megaprem Bot is ready. Type m!help to see commands.")
		}
	}
}

func helpText() string {
	return `Megaprem Bot Help
m!help: displays this message`
}
