package tnyuri

import (
	"net/url"
	"strconv"
)

func isURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func exists(user string, id int) bool {
	return Select[int]("SELECT count(*) FROM url WHERE user='"+user+"' AND id='"+strconv.Itoa(id)+"'") > 0
}

func (b *Bot) Add(user string, url string) {
	if isURL(url) {

		if Select[int]("SELECT COUNT(*) FROM url WHERE user='"+user+"'") >= config.Maxperuser {
			b.Send("Maximum URL for your users has been reach")
		} else {
			if Select[int]("SELECT COUNT(*) FROM url WHERE user='"+user+"' AND url='"+url+"'") > 0 {
				b.Send("Your link is existing")
			} else {
				var nlink *URL = new(URL)
				nlink.User = user
				nlink.Url = url
				var id int = nlink.Save()
				if id >= 0 {
					b.Send("Your link was created: " + config.Domain + nlink.Short)
					b.Send("`" + strconv.Itoa(id) + ", " + nlink.Url + ", " + config.Domain + nlink.Short + "`")

				} else {
					b.Send("Cant create this link")
				}
			}
		}
	} else {
		b.Send("Wrong url Parameter")
	}
}

func (b *Bot) Delete(user string, id int) {
	if exists(user, id) {
		Delete[URL](user, id)
		b.Send("ID: " + strconv.Itoa(id) + " deleted!")
	}
}

func (b *Bot) Stats(user string, id int) {

	if exists(user, id) {
		var url URL = Select[URL]("SELECT *FROM url WHERE user='" + user + "' AND id='" + strconv.Itoa(id) + "'")

		var stats Stats = url.Stats()

		b.Send("`" + strconv.Itoa(url.Id) + ", " + url.Url + ", " + config.Domain + url.Short + ", " + strconv.Itoa(stats.Counter) + "`")
	}
}
