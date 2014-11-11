package main

import (
	"bufio"
	"github.com/merlinran/systray"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		println("usage: example icon-path client-path")
		return
	}

	menu := []systray.MenuItem{
		{"Show Info", "info"},
		{"Quit", "quit"},
	}

	tray := systray.New(os.Args[1], os.Args[2], 6333)

	quit := func() {
		err := tray.Stop()
		if err != nil {
			println(err.Error())
		}
		os.Exit(0)
	}

	tray.OnClick(func() {
		println("clicked")
	})
	tray.OnAction(func(actionName string) {
		switch actionName {
		case "info":
			println("Show Info menu clicked")
		case "quit":
			quit()
		}
	})
	err := tray.Show("idle.ico", "Test systray", menu[:])
	if err != nil {
		println(err.Error())
	}

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			println("Input icon file name:")
			print(">> ")
			data, _, _ := reader.ReadLine()
			line := string(data)
			if len(line) == 0 {
				break
			}
			err := tray.Show(line, line, menu[:])
			if err != nil {
				println(err.Error())
			}
		}
		quit()
	}()

	err = tray.Run()
	if err != nil {
		println(err.Error())
	}
}
