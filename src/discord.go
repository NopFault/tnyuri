package tnyuri

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

/**
 * TODO:
 *
 * Fix this global madness and make some logic on top of it!!!
 * */

var Botas *discordgo.User
var Session *discordgo.Session

func BotStart() {
	session, err := discordgo.New("Bot " + config.Token)
	Session = session

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot, err := Session.User("@me")
	Botas = bot

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Session.AddHandler(MessageHandler)
	err = Session.Open()
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

func NotifyUser(userId string, message string) bool {
	if Session != nil {
		channel, _ := Session.UserChannelCreate(userId)
		Session.ChannelMessageSend(channel.ID, message)
		return true
	}
	return false
}

func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == Botas.ID {
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
