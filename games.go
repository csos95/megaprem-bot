package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"strings"
)

func roll(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	switch len(args) {
	case 0:
		sendMessage(s, m.ChannelID, fmt.Sprintf("You rolled a %d", rand.Int31n(5)+1))
	case 1:
		rolls, err := strconv.Atoi(args[0])
		if err != nil {
			sendMessage(s, m.ChannelID, "Please use an integer for the number of dice to roll.")
			return
		}
		if rolls > 100 {
			sendMessage(s, m.ChannelID, "The max number of rolls is 100.")
			return
		}
		if rolls < 0 {
			sendMessage(s, m.ChannelID, "Negative numbers of rolls are not allowed.")
			return
		}

		results := make([]string, rolls)
		for i := range results {
			results[i] = strconv.Itoa(int(rand.Int31n(5) + 1))
		}
		sendMessage(s, m.ChannelID, fmt.Sprintf("Your results: %s", strings.Join(results, " ")))
	case 2:
		sides, err := strconv.Atoi(args[0])
		if err != nil {
			sendMessage(s, m.ChannelID, "Please use an integer for the number of dice to roll.")
			return
		}
		if sides > 20 {
			sendMessage(s, m.ChannelID, "The max number of sides is 20.")
			return
		}
		if sides < 4 {
			sendMessage(s, m.ChannelID, "The minumum number of sides is 4.")
			return
		}
		rolls, err := strconv.Atoi(args[1])
		if err != nil {
			sendMessage(s, m.ChannelID, "Please use an integer for the number of dice to roll.")
			return
		}
		if rolls > 100 {
			sendMessage(s, m.ChannelID, "The max number of rolls is 100.")
			return
		}
		if rolls < 0 {
			sendMessage(s, m.ChannelID, "Negative numbers of rolls are not allowed.")
			return
		}

		results := make([]string, rolls)
		for i := range results {
			results[i] = strconv.Itoa(int(rand.Int31n(int32(sides-1)) + 1))
		}
		sendMessage(s, m.ChannelID, fmt.Sprintf("Your results: %s", strings.Join(results, " ")))
	default:
		sendMessage(s, m.ChannelID, "Too many arguments.")
	}
}

func poll(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	sendMessage(s, m.ChannelID, "Polls have not been implemented.")
	return
}
