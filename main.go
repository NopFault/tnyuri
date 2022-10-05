package main

import (
	"fmt"
	tnyuri "tnyuri/src"
)

func main() {

	// Config test
	// var conf tnyuri.Config = tnyuri.GetConfig()
	// fmt.Println(conf.Token)

	// BOT Start
	tnyuri.BotStart()
	fmt.Println("After BOT INIT")
	// DB Init test
	tnyuri.Init()
	fmt.Println("After DB INIT")
	// Start WEB server
	tnyuri.InitWeb()
	fmt.Println("After WEB INIT")

	// DB Insert func Test
	// tnyuri.Insert("insert into url ('url', 'short', 'user', 'status') VALUES ('http://x90.lt','abc','nopfault',true)")
	// fmt.Println("Inserted")

	// count := tnyuri.Select[int]("select count(*) from url")
	// fmt.Println("Kiekis: " + strconv.Itoa(count))

	// user := tnyuri.Select[string]("select user from url where id=1")
	// fmt.Println("Username: " + user)

	// d := tnyuri.By[tnyuri.Stats]("url_id", "1")
	// fmt.Println(d)
	//
	// var nlink *tnyuri.URL = new(tnyuri.URL)
	// nlink.User = "Paul"
	// nlink.Url = "https://x90.lt"

	// nlink.Save()
	// //
	// fmt.Println(tnyuri.By[tnyuri.URL]("user", "Paul"))
	//

	// d := tnyuri.By[tnyuri.URL]("id", "2")
	//
	// sd := d[0].Stats()
	// sd.Increase()
	//
	// fmt.Println(sd)

}
