package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

func init() {
	commands["ui"] = testUI
}

var root *tview.Flex
var app *tview.Application
var focusChain []tview.Primitive
var focusIndex = 0

var debug *tview.TextView

func testUI(args []string) {
	app = tview.NewApplication()

	root = tview.NewFlex().SetDirection(tview.FlexRow)
	app.SetRoot(root, true)

	cmsg := tview.NewInputField()
	cmsg.SetBorder(true).SetTitle("Commit Message")
	cmsg.SetDoneFunc(func(key tcell.Key) {
		moveFocus(1)
	})

	mode := tview.NewList()
	mode.SetBorder(true).SetTitle("Mode")
	mode.ShowSecondaryText(false).
		AddItem("Full", "", 'f', nil).
		AddItem("Partial", "", 'p', nil)

	debug = tview.NewTextView()

	root.AddItem(cmsg, 3, 0, false)
	root.AddItem(mode, 4, 0, false)
	root.AddItem(debug, 0, 1, false)

	focusChain = []tview.Primitive{cmsg, mode}
	app.SetFocus(focusChain[focusIndex])

	err := app.SetInputCapture(handleKey).Run()
	CheckIfError(err)
}

func moveFocus(delta int) {
	focusIndex += delta
	if focusIndex < 0 {
		focusIndex = 0
	}
	if focusIndex >= len(focusChain) {
		focusIndex = len(focusChain) - 1
	}
	app.SetFocus(focusChain[focusIndex])
	app.Draw()
}

func handleKey(key *tcell.EventKey) *tcell.EventKey {
	debug.SetText(fmt.Sprintf("name:%b  mods:%b", key.Name(), key.Modifiers()))

	switch key.Key() {
	case tcell.KeyTab:
		moveFocus(1)
		return nil
	case tcell.KeyBacktab:
		moveFocus(-1)
		return nil
	case tcell.KeyCtrlJ:
		debug.SetText("Control Enter Detected")
	}
	return key
}
