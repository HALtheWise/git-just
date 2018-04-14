package main // import "github.com/HALtheWise/git-just"

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

func main() {
	_ = git.Added
	if len(os.Args) <= 1 {
		fmt.Println("TODO: usage menu")
		return
	}
	command := os.Args[1]
	f, ok := commands[command]
	if !ok {
		fmt.Println("TODO: typo menu")
		return
	}

	f(os.Args[2:])
}

type CommandFunc func([]string)

var commands map[string]CommandFunc

func init() {
	commands = make(map[string]CommandFunc)
	commands["save"] = cmdSave
	commands["s"] = commands["save"]
}

func cmdSave(args []string) {
	fmt.Println("Saving")
}
