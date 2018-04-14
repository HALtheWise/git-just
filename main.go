package main // import "github.com/HALtheWise/git-just"

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

func main() {
	_ = git.Added
	fmt.Println("Hello World", os.Args[1:])
}
