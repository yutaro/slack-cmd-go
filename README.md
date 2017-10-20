# slack-cmd-go

Making slack-bot perseing your messages like cli commands.

```go
package main

import(
    "github.com/yutaro/slack-cmd-go"
)

func main(){
    bot := scmd.New("--- YOUR API KEY ---")
    calc := bot.NewCmd("calc")

    calc.Action("sum", "Add two numbers.",
        func(c scmd.Context){
        args := c.GetArgs()
        x, _ := strconv.Atoi(args[0])
        y, _ := strconv.Atoi(args[1])
        c.SendMessage(fmt.Sprintf("The result is : %d", x + y))
    })

    calc.Action("fib", "Show fibonacci numbers." ,
        func(c scmd.Context){
        args := c.GetArgs()
        x, _ := strconv.ParseInt(args[0])

        nums := make(string, x)
        f := fibbonacci()
        for i := 0; i < x; i++{
            nums[x] = strconv.Itoa(f())
        }

        c.SendMessage(fmt.Sprintf("The result is : %s", strings.Join(nums, " ")))
    })

    bot.Start();
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