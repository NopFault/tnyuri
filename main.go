package main

import (
	"fmt"
	tnyuri "tnyuri/src"
)

func main() {

	// BOT Start
	tnyuri.BotStart()
	fmt.Println("After BOT INIT")
	// DB Init test
	tnyuri.Init()
	fmt.Println("After DB INIT")
	// Start WEB server
	tnyuri.InitWeb()
	fmt.Println("After WEB INIT")

}
