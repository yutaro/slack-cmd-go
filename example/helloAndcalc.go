package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yutaro/slack-cmd-go"
)

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)

	bot.OneCmd("hello", "greeting",
		func(c *scmd.Context) {
			args := c.GetArgs()
			if len(args) == 0 {
				c.SendMessage("Hello!")
				return
			}
			c.SendMessage(fmt.Sprintf("Hello %s!", strings.Join(args, " ")))
		})

	calc := bot.NewCmds("calc")

	calc.Cmd("sum", "Add two numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x, _ := strconv.Atoi(args[0])
			y, _ := strconv.Atoi(args[1])
			c.SendMessage(fmt.Sprintf("The result is : %d", x+y))
		})

	calc.Cmd("fib", "Show fibonacci numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x, _ := strconv.Atoi(args[0])

			nums := make([]string, x)
			f := fibonacci()
			for i := 0; i < x; i++ {
				nums[x] = strconv.Itoa(f())
			}

			c.SendMessage(fmt.Sprintf("The result is : %s", strings.Join(nums, " ")))
		})

	bot.Start()
}

func fibonacci() func() int {
	fib1 := 0
	fib2 := 1
	return func() int {
		fib := fib1 + fib2
		fib1 = fib2
		fib2 = fib
		return fib1
	}
}
