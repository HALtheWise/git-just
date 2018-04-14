package main // import "github.com/HALtheWise/git-just"

import (
	"fmt"
	"os"

	"time"

	"strings"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
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

	if len(args) == 0 {
		fmt.Println("Please give me a commit message! \n\tTODO: make a field")
		return
	}
	message := strings.Join(args, " ")

	dir, err := os.Getwd()
	CheckIfError(err)
	Info("Working in directory %s", dir)

	// Opens an already existent repository.
	r, err := git.PlainOpen(dir)
	CheckIfError(err)

	var w *git.Worktree

	w, err = r.Worktree()
	CheckIfError(err)

	// Adds the new file to the staging area.
	Info("git add main.go")
	_, err = w.Add("main.go")
	CheckIfError(err)

	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)
	fmt.Println(status)

	// Commits the current staging are to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit.
	Info("git commit -m \"%s\"", message)
	commit, err := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)

	Info("git push origin")
	r.Push(&git.PushOptions{})
}
