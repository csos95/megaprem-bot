package main

import "github.com/bwmarrin/discordgo"

type Command struct {
	name        string
	description string
	arguments   []string
	function    func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

func NewCommand(name, description string, arguments []string, function func(*discordgo.Session, *discordgo.MessageCreate, []string)) Command {
	return Command{name: name, description: description, arguments: arguments, function: function}
}
