# slack-cmd-go

Making slack-bot perseing your messages like cli commands.

```go
package main

import(
    "github.com/yutaro/slack-cmd-go"
)

func main(){
    bot := scmd.New("--- YOUR API KEY ---")
    calc := api.NewCmd("calc")

    calc.GET("sum", func(c scmd.Context){
        args := c.GetArgs()
        x, _ = strconv.ParseInt(args[0])
        y, _ = strconv.ParseInt(args[1])
        c.SendMessage(fmt.Sprintf("The result is : %d", x + y))
    })

    bot.Start();
}


```