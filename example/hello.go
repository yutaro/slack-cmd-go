package main

import (
	"fmt"
	"strings"

	"github.com/yutaro/slack-cmd-go"
)

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)

	// just one phrase command.
	// hello => Hello!
	// hello yutaro => Hello yutaro!
	bot.OneCmd("hello", []string{"greeting"},
		func(c *scmd.Context) {
			args := c.GetArgs()
			if len(args) == 0 {
				c.SendMessage("Hello!")
				return
			}
			c.SendMessage(fmt.Sprintf("Hello %s!", strings.Join(args, " ")))
		})

	bot.Start()
}
