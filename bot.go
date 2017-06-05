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
	config          *Config
}

func CreateBot(config *Config) error {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return err
	}

	bot = &Bot{
		Session:         dg,
		commands:        make([]Command, 0),
		messageLifetime: time.Second * time.Duration(config.MessageLifetime),
		config:          config,
	}

	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
	bot.AddHandler(guildCreate)

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

	if strings.HasPrefix(m.Content, bot.config.Prefix) {
		parts := strings.Fields(m.Content[len(bot.config.Prefix):])
		if len(parts) != 0 {
			for _, command := range bot.commands {
				if parts[0] == command.name {
					command.function(s, m, parts[1:])

					if bot.messageLifetime != 0 {
						lifetimeChan := time.After(bot.messageLifetime)
						go func() {
							<-lifetimeChan
							deleteMessage(s, m.ChannelID, m.Message)
						}()
					}
					break
				}
			}
		}
	}
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	//if event.Guild.Unavailable {
	//	return
	//}
	//
	//for _, channel := range event.Guild.Channels {
	//	if channel.ID == event.Guild.ID && channel.ID != "96081945389182976" {
	//		sendMessage(s, channel.ID, fmt.Sprintf("Megaprem Bot is ready. Type %shelp to see commands.", bot.config.Prefix))
	//	}
	//}
}

func addCommands() {
	commands := []Command{
		NewCommand("help", "displays this message", []string{}, help),
		NewCommand("roll", "roll dice", []string{"[number]", "[sides] [number]"}, roll),
		NewCommand("messageLifetime", "set message lifetime", []string{"[seconds]"}, messageLifetime),
		NewCommand("setPrefix", "set the command prefix", []string{"[prefix]"}, prefix),
		NewCommand("imgur", "search imgur", []string{"[query]"}, imgur),
		NewCommand("giphy", "search giphy", []string{"[query]"}, giphy),
		NewCommand("lmgtfy", "make a let me google that for you link", []string{"[query]"}, lmgtfy),
		NewCommand("google", "search google images", []string{"[query]"}, google),
		NewCommand("poll", "make a poll", []string{"name:[name] duration:[duration in seconds] options:[comma separaed options]"}, poll),
		NewCommand("setStatus", "set the bot status", []string{"[status]"}, status),
	}
	bot.commands = append(bot.commands, commands...)
}

func createHelp() {
	help := fmt.Sprintf("Megaprem Bot Help\n\n\tcommand: description\n\t\targs\n\n")

	for _, command := range bot.commands {
		help += fmt.Sprintf("\t%s: %s\n\t\t%s\n", command.name, command.description, strings.Join(command.arguments, "\n\t\t"))
	}
	bot.help = help
}
