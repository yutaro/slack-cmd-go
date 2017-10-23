package scmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getNewHelp(name string) *Cmd {
	cmd := &Cmd{
		name:    name,
		label:   "help",
		explain: "show usage.",
		run: func(c *Context) {
			c.SendMessage(c.GetHelpMes(name))
		},
	}

	return cmd
}

func (c *Context) GetHelpMes(name string) string {
	commands := c.bot.Cmds[name]
	labels := make([]string, 0, len(commands))
	for label := range commands {
		labels = append(labels, label)
	}
	sort.Strings(labels)

	maxLen := 0
	for _, u := range commands {
		if len(u.label) > maxLen {
			maxLen = len(u.label)
		}
	}

	mes := fmt.Sprintf("```\n%s", name)
	spaces := strings.Repeat(" ", len(name))
	flag := true

	form1 := " %-" + strconv.Itoa(maxLen) + "s : %s"
	form2 := "\n%s" + form1
	for _, label := range labels {
		u := commands[label]
		if flag {
			mes += fmt.Sprintf(form1, u.label, u.explain)
			flag = false
			continue
		}

		mes += fmt.Sprintf(form2, spaces, u.label, u.explain)
	}
	mes += "\n```"

	return mes
}
