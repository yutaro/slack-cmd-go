package main

import (
	"github.com/yutaro/slack-cmd-go"
)

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)

	bot.OneCmd("image", "upload image", func(c *scmd.Context) {
		c.SendFile("myaox.jpg")
	})

	bot.Start()
}
