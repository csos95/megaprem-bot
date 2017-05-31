package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var bot *Bot

type Bot struct {
	*discordgo.Session
	commands        []Command
	help            string
	messageLifetime time.Duration
	prefix          string
}

func CreateBot(config *Config) error {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return err
	}

	bot = &Bot{Session: dg, commands: make([]Command, 0), prefix: "m!"}

	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
	bot.AddHandler(guildCreate)

	bot.messageLifetime = time.Second * 10

	addCommands()

	createHelp()

	return nil
}

func run() error {
	err := bot.Open()
	if err != nil {
		return err
	}

	fmt.Println("Megaprem Bot is now running. Press CRTL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
	return nil
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "m!help")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, bot.prefix) {
		parts := strings.Fields(strings.TrimLeft(m.Content, bot.prefix))
		for _, command := range bot.commands {
			if parts[0] == command.name {
				command.function(s, m, parts[1:])
				break
			}
		}
	}

	if bot.messageLifetime != 0 {
		lifetimeChan := time.After(bot.messageLifetime)
		go func() {
			<-lifetimeChan
			deleteMessage(s, m.ChannelID, m.Message)
		}()
	}
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID && channel.ID != "96081945389182976" {
			sendMessage(s, channel.ID, "Megaprem Bot is ready. Type m!help to see commands.")
		}
	}
}

func addCommands() {
	commands := []Command{
		NewCommand("help", "displays this message", []string{}, helpText),
		NewCommand("roll", "roll dice", []string{"[number]", "[sides] [number]"}, roll),
		NewCommand("messageLifetime", "set message lifetime", []string{"[seconds]"}, setMessageLifetime),
		NewCommand("setPrefix", "set the command prefix", []string{"[prefix]"}, setPrefix),
	}
	bot.commands = append(bot.commands, commands...)
}

func createHelp() {
	help := `Megaprem Bot Help

	command: description
		args

`
	for _, command := range bot.commands {
		help += fmt.Sprintf("\t%s: %s\n\t\t%s\n", command.name, command.description, strings.Join(command.arguments, "\n\t\t"))
	}
	bot.help = help
}
