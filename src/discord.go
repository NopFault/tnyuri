package tnyuri

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var bot *discordgo.User
var session *discordgo.Session

func BotStart() {
	session, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot, err = session.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	session.AddHandler(MessageHandler)

	err = session.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
func validUrl(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
func sendMessage(msg string, message *discordgo.MessageCreate, session *discordgo.Session) {
	session.ChannelMessageSend(message.ChannelID, msg)
}
func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == bot.ID {
		return
	}
	var data []string = strings.Split(message.Content, " ")

	if data[0] == "tnyuri" && len(data) >= 2 {
		var bot *Bot = new(Bot)

		bot.Session = session
		bot.Message = message
		bot.Handle(data)
	}
}
