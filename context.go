package scmd

import (
	"fmt"

	"github.com/nlopes/slack"
)

type Context struct {
	rtm     *slack.RTM
	ev      *slack.MessageEvent
	bot     *Bot
	options map[string]string
	flags   map[string]bool
	rawArgs []string
	args    []string
}

func (c *Context) SendMessage(mes string) {
	c.rtm.SendMessage(c.rtm.NewOutgoingMessage(mes, c.ev.Channel))
}

func (c *Context) SendFile(mes string) {
	file, err := c.rtm.UploadFile(&slack.params{
		Title: mes,
		File:  mes,
	})

	if err != nil {
		return
	}
	fmt.Println(file)
}

type argStr string

func (c *Context) GetArgs() []string {
	return c.args
}

func (c *Context) GetFlags() map[string]bool {
	return c.flags
}

func (c *Context) GetOptions() map[string]string {
	return c.options
}
