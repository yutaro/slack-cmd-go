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
			flags := c.GetFlags()
			options := c.GetOptions()

			fmt.Println(flags)
			fmt.Println(options)
		})

	bot.Start()
}
