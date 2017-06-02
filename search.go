package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"strings"
)

type ImgurImage struct {
	Link string `json:"link"`
}

type ImgurResponse struct {
	Data    []ImgurImage `json:"data"`
	Success bool         `json:"success"`
}

func imgur(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		sendMessage(s, m.ChannelID, "Not enough arguments.")
		return
	}

	url := fmt.Sprintf("https://api.imgur.com/3/gallery/search/0?q=%s", strings.Join(args, ","))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	request.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", bot.config.ImgurID))
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	imgurResponse := ImgurResponse{}

	err = json.NewDecoder(response.Body).Decode(&imgurResponse)
	if err != nil {
		log.Println(err)
		return
	}

	switch len(imgurResponse.Data) {
	case 0:
		sendMessage(s, m.ChannelID, "No results.")
	case 1:
		sendRawMessage(s, m.ChannelID, imgurResponse.Data[0].Link)
	case 2:
		sendRawMessage(s, m.ChannelID, imgurResponse.Data[0].Link+"\n"+imgurResponse.Data[1].Link)
	default:
		sendRawMessage(s, m.ChannelID, imgurResponse.Data[0].Link+"\n"+imgurResponse.Data[1].Link+"\n"+imgurResponse.Data[2].Link)
	}
}

type GiphyGif struct {
	Link string `json:"embed_url"`
}

type GiphyResponse struct {
	Data []GiphyGif `json:"data"`
}

func giphy(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		sendMessage(s, m.ChannelID, "Not enough arguments.")
		return
	}

	url := fmt.Sprintf("http://api.giphy.com/v1/gifs/search?q=%s&api_key=%s", strings.Join(args, "+"), bot.config.GiphyKey)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	giphyResponse := GiphyResponse{}

	err = json.NewDecoder(response.Body).Decode(&giphyResponse)
	if err != nil {
		log.Println(err)
		return
	}

	switch len(giphyResponse.Data) {
	case 0:
		sendMessage(s, m.ChannelID, "No results.")
	case 1:
		sendRawMessage(s, m.ChannelID, giphyResponse.Data[0].Link)
	case 2:
		sendRawMessage(s, m.ChannelID, giphyResponse.Data[0].Link+"\n"+giphyResponse.Data[1].Link)
	default:
		sendRawMessage(s, m.ChannelID, giphyResponse.Data[0].Link+"\n"+giphyResponse.Data[1].Link+"\n"+giphyResponse.Data[2].Link)
	}
}

func lmgtfy(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		sendMessage(s, m.ChannelID, "Not enough arguments.")
		return
	}

	sendRawMessage(s, m.ChannelID, fmt.Sprintf("http://lmgtfy.com/?q=%s", strings.Join(args, "+")))
}

type GoogleImage struct {
	Link string `json:"link"`
}

type GoogleResult struct {
	Data []GoogleImage `json:"items"`
}

func google(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		sendMessage(s, m.ChannelID, "Not enough arguments.")
		return
	}

	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&searchType=image&q=%s", bot.config.GoogleApi, bot.config.GoogleSearch, strings.Join(args, "+"))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	googleResponse := GoogleResult{}

	err = json.NewDecoder(response.Body).Decode(&googleResponse)
	if err != nil {
		log.Println(err)
		return
	}

	switch len(googleResponse.Data) {
	case 0:
		sendMessage(s, m.ChannelID, "No results.")
	case 1:
		sendRawMessage(s, m.ChannelID, googleResponse.Data[0].Link)
	case 2:
		sendRawMessage(s, m.ChannelID, googleResponse.Data[0].Link+"\n"+googleResponse.Data[1].Link)
	default:
		sendRawMessage(s, m.ChannelID, googleResponse.Data[0].Link+"\n"+googleResponse.Data[1].Link+"\n"+googleResponse.Data[2].Link)
	}
}
