package megaprem_bot

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os/signal"
	"syscall"
	"os"
)

type bot struct {
	*discordgo.Session
}

func NewBot(config *Config) (*bot, error) {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	bot := &bot{dg}

	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
	bot.AddHandler(guildCreate)

	return bot, nil
}

func (b *bot) run() error {
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

}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Megaprem Bot is ready. Type m!help to see commands.")
		}
	}
}
