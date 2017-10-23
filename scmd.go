package scmd

import (
	"fmt"
	"strings"

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
		api:  api,
		rtm:  rtm,
		Cmds: make(map[string]map[string]*Cmd),
	}
}

func (b *Bot) NewCmds(name string) *CmdGroup {
	b.Cmds[name] = make(map[string]*Cmd)
	return &CmdGroup{
		name: name,
		Bot:  *b,
	}
}

func (b *Bot) OneCmd(name, explain string, callback func(*Context)) {
	b.Cmds[name] = make(map[string]*Cmd)
	b.Cmds[name][" "] = &Cmd{
		name:    name,
		label:   " ",
		explain: explain,
		run:     callback,
	}
}

func (g *CmdGroup) Cmd(label, explain string, callback func(*Context)) {
	g.Bot.Cmds[g.name][label] = &Cmd{
		name:    g.name,
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
			case *slack.RTMError:
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
	rawArgs := strings.Split(msg, " ")
	args := make([]string, 0)

	c := &Context{
		rtm:     b.rtm,
		ev:      ev,
		rawArgs: rawArgs,
		options: make(map[string]string),
		flags:   make(map[string]bool),
	}

	group, ok := b.Cmds[rawArgs[0]]
	if !ok {
		//fmt.Printf("Not exist command group %s", args[0])
		return
	}

	for _, a := range rawArgs {
		if a == "" {
			continue
		}
		if len(a) < 2 {
			args = append(args, a)
			continue
		}

		var dh bool
		if a[0:2] == "--" {
			a = a[2:]
			dh = true
		} else if a[0] == '-' {
			a = a[1:]
			dh = false
		} else {
			args = append(args, a)
			continue
		}

		if len(a) < 1 {
			args = append(args, a)
			continue
		}

		if strings.Contains(a, "=") {
			vals := strings.Split(a, "=")
			label := vals[0]
			val := strings.Join(vals[1:], "=")
			c.options[label] = val
		} else {
			if dh {
				c.flags[a] = true
			} else {
				for _, label := range a {
					c.flags[string(label)] = true
				}
			}
		}
	}

	if cmd, ok := group[" "]; ok && len(args) == 1 {
		c.args = args[1:]
		cmd.run(c)
		return
	}

	if cmd, ok := group[args[1]]; ok {
		c.args = args[2:]
		cmd.run(c)
		return
	} else if cmd, ok := group[" "]; ok {
		c.args = args[1:]
		cmd.run(c)
		return
	} else {
		fmt.Printf("Not Exist command: %s %s\n", args[0], args[1])
		return
	}
}
