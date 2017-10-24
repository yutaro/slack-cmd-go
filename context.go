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

func (c *Context) SendFile(mes string) *slack.File {
	file, err := c.rtm.UploadFile(slack.FileUploadParameters{
		Title:    mes,
		File:     mes,
		Channels: []string{c.ev.Channel},
	})

	if err != nil {
		return nil
	}
	//fmt.Printf("%+v\n", file)
	return file
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
