package scmd

import "github.com/nlopes/slack"

type Context struct {
	rtm     *slack.RTM
	ev      *slack.MessageEvent
	options map[string]string
	flags   map[string]bool
	rawArgs []string
	args    []string
}

func (c *Context) SendMessage(mes string) {
	c.rtm.SendMessage(c.rtm.NewOutgoingMessage(mes, c.ev.Channel))
}

func (c *Context) GetArgs() []string {
	return c.args
}

type Option struct {
	name    string
	isFlag  bool
	flagVal bool

	value string
}
