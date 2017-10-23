# slack-cmd-go

Making slack-bot perseing your messages like cli commands.

## Usage

### install
```sh
go get -u github.com/yutaro/slack-cmd-go
```


## Examples


## Hello World

```go
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

	// just one phrase command.
	// hello => Hello!
	// hello yutaro => Hello yutaro!
	bot.OneCmd("hello", "greeting",
		func(c *scmd.Context) {
			args := c.GetArgs()
			if len(args) == 0 {
				c.SendMessage("Hello!")
				return
			}
			c.SendMessage(fmt.Sprintf("Hello %s!", strings.Join(args, " ")))
		})

	bot.Start()
}

```

<img src="https://i.gyazo.com/276c661ed615951e83f11907b54d8e86.png" width="500">

## Calc

```go
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

	// two phrase command
	calc := bot.NewCmds("calc")

	// calc sum 2 3 => The result is : 5
	calc.Cmd("sum", "Add two numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x, _ := strconv.Atoi(args[0])
			y, _ := strconv.Atoi(args[1])
			c.SendMessage(fmt.Sprintf("The result is : %d", x+y))
		})

	// calc sub 5 10 => The result is : -5
	calc.Cmd("sub", "Sub two numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x, _ := strconv.Atoi(args[0])
			y, _ := strconv.Atoi(args[1])
			c.SendMessage(fmt.Sprintf("The result is : %d", x-y))
		})

	// calc fib 5 => The result is : 1 1 2 3 5
	calc.Cmd("fib", "Show fibonacci numbers.",
		func(c *scmd.Context) {
			args := c.GetArgs()
			x, _ := strconv.Atoi(args[0])

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
```

<img src="https://i.gyazo.com/6c3bd2b40cd1cd7a7e792f23982411ed.png" width="500">

## Options

```go
package main

import (
	"fmt"

	"github.com/yutaro/slack-cmd-go"
)

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)

	bot.OneCmd("test", "test options and flags",
		func(c *scmd.Context) {
			args := c.GetArgs()
			flags := c.GetFlags()
			options := c.GetOptions()

			c.SendMessage("--- test ---")
			c.SendMessage(fmt.Sprintf("options : %v , flags : %v , args : %v \n", options, flags, args))
		})

	bot.Start()
}
```

<img src="https://i.gyazo.com/6d8c16edd3680bb394599ff34cc246d7.png" width="500">

## Library
<https://github.com/nlopes/slack>    
<https://github.com/BurntSushi/toml>
