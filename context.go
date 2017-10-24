package scmd

import (
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
	c.rtm.UploadFile(slack.FileUploadParameters{
		Title:    mes,
		File:     mes,
		Channels: []string{c.ev.Channel},
	})

	//fmt.Printf("%+v\n", file)
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
