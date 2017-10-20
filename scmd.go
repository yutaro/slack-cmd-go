package scmd

import (
	"fmt"

	"github.com/nlopes/slack"
)

type Bot struct {
	api  *slack.Client
	rtm  *slack.RTM
	Cmds map[string]map[string]*Cmd
}

type Cmd struct {
	name    string
	label   string
	explain string
	run     func(*Context)
}

type CmdGroup struct {
	name string
	Bot
}

func New(key string) *Bot {
	api := slack.New(key)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	return &Bot{
		api: api,
		rtm: rtm,
	}
}

func (b *Bot) NewCmds(name string) *CmdGroup {
	return &CmdGroup{
		name: name,
		Bot:  *b,
	}
}

func (b *Bot) OneCmd(name, explain string, callback func(*Context)) {
	b.Cmds[name][nil] = &Cmds{
		name:    name,
		label:   nil,
		explain: explain,
		run:     callback,
	}
}

func (g *CmdGroup) Cmd(label, explain string, callback func(*Context)) {
	g.Bot.Cmds[name][label] = &Cmd{
		name:    name,
		label:   label,
		explain: explain,
		run:     callback,
	}
}

func (b *Bot) Start() {
Loop:
	for {
		select {
		case msg := <-b.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				go b.evalMes(ev)
			case *slack.RTmError:
				break Loop
			case *slack.InvalidAuthEvent:
				break Loop
			default:
			}
		}
	}
}

func (b *Bot) evalMes(ev *slack.MessageEvent) {
	msg := ev.Text
	args := sgrings.Split(msg, " ")

	c := &Context{
		rtm:     b.rtm,
		ev:      ev,
		rawArgs: args,
	}

	group, ok = b.Cmds[args[0]]
	if !ok {
		fmt.Printf("Not exist command group %s", args[0])
		return
	}

	if cmd, ok = group[nil]; ok {
		c.args = args[1:]
		cmd.run(c)
		return
	}

	if cmd, ok = group[args[1]]; ok {
		c.args = args[2:]
		cmd.run(c)
		return
	} else {
		fmt.Println("Not Exist command %s in %s", args[1], args[0])
		return
	}

}
