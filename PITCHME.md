# slack cli化

---

### 今までの経緯
- CLIっぽくできると履歴が見れる
- 他の人の利用状況を管理できる
- 単純に楽なケースもある

など便利な部分があった

---

#### 色々な分野に応用するのがちょっとめんどくさい

---

そこでライブラリを作った


<https://github.com/yutaro/slack-cmd-go>

---
hello => helloの命令
```go
bot.OneCmd("hello", "greeting",
	func(c *scmd.Context) {
		args := c.GetArgs()
		if len(args) == 0 {
			c.SendMessage("Hello!")
			return
		}
		c.SendMessage(fmt.Sprintf("Hello %s!", strings.Join(args, " ")))
	})
```

---

calc sum 2 3 => 5の命令

```go
calc := bot.NewCmds("calc")
calc.Cmd("sum", "Add two numbers.",
	func(c *scmd.Context) {
		args := c.GetArgs()
		x, _ := strconv.Atoi(args[0])
		y, _ := strconv.Atoi(args[1])
		c.SendMessage(fmt.Sprintf("The result is : %d", x+y))
	})
```

---

## かなり簡単にかけるようになった！

---

# 終わり
