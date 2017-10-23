package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/spf13/cast"

	"github.com/yutaro/slack-cmd-go"
)

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)

	// two phrase command
	calc := bot.NewCmdGroup("calc")

	// calc sum 2 3 => The result is : 5
	calc.Cmd("sum", "Add two numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x := ToInt(args[0])
			y := ToInt(args[1])
			c.SendMessage(fmt.Sprintf("The result is : %d", x+y))
		})

	// calc sub 5 10 => The result is : -5
	calc.Cmd("sub", "Sub two numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x := ToInt(args[0])
			y := ToInt(args[1])
			c.SendMessage(fmt.Sprintf("The result is : %d", x-y))
		})

	// calc fib 5 => The result is : 1 1 2 3 5
	calc.Cmd("fib", "Show fibonacci numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x := ToInt(args[0])

			nums := make([]string, x)
			f := fibonacci()
			for i := 0; i < x; i++ {
				nums[i] = strconv.Itoa(f())
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
