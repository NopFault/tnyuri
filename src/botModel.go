package tnyuri

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
}

func (b *Bot) Send(msg string) {
	b.Session.ChannelMessageSend(b.Message.ChannelID, msg)
}

func (b *Bot) Handle(data []string) {

	var username string = b.Message.Author.Username
	var action string = data[1]

	if action == "add" {
		if len(data) > 2 {
			// botactions
			b.Add(username, data[2])
		}

	} else if action == "delete" {
		if len(data) > 2 {
			// botactions
			id, err := strconv.Atoi(data[2])
			if err != nil {
				b.Send("Wrong ID Parameter")
			}
			b.Delete(username, id)
		}
	} else if action == "stats" {
		if len(data) > 2 {
			// botactions
			id, err := strconv.Atoi(data[2])
			if err != nil {
				b.Send("Wrong ID Parameter")
			}
			b.Stats(username, id)
		}
	} else if action == "list" {
		var urls []URL = RowsBy[URL]("url", "user", username)
		var strList string = "```"

		if len(urls) > 0 {
			for _, url := range urls {
				var stats Stats = url.Stats()
				strList += strconv.Itoa(url.Id) + ", " + url.User + ", " + url.Url + ", " + config.Domain + url.Short + ", " + strconv.Itoa(stats.Counter) + "\n"
			}
			strList += "```"
			b.Send(strList)
		} else {
			b.Send("You dont have any links created")
		}
	}
}
