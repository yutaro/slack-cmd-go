package main

import (
	"fmt"

	"github.com/yutaro/slack-cmd-go"
)

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)

	bot.OneCmd("test", "test options and flags",
		func(c *scmd.Context) {
			args := c.GetArgs()
			flags := c.GetFlags()
			options := c.GetOptions()

			fmt.Println("--- test ---")
			fmt.Printf("options : %v , flags : %v , args : %v \n", options, flags, args)
		})

	bot.Start()
}
