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
		explain: []string{"show usage."},
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
	cmdSpace := strings.Repeat(" ", len(name))
	labelSpace := strings.Repeat(" ", maxLen)
	flag := true

	form1 := " %-" + strconv.Itoa(maxLen) + "s : %s"
	form2 := "\n%s" + form1
	form3 := "\n%s %s   %s"

	for _, label := range labels {
		u := commands[label]
		for i, exp := range u.explain {
			if flag {
				mes += fmt.Sprintf(form1, u.label, exp)
				flag = false
				continue
			}

			if i == 0 {
				mes += fmt.Sprintf(form2, cmdSpace, u.label, exp)
			} else {
				mes += fmt.Sprintf(form3, cmdSpace, labelSpace, exp)
			}
		}
	}
	mes += "\n```"

	return mes
}
