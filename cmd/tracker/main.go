package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	version = "1.0.0"

	fileName = "/tmp/time-tracker.log"

	usage = "usage:\n\ttt [start|stop|status]"
)

type application struct {
	action  string
	storage struct {
		name string
		file *os.File
	}
}

func main() {
	var app application

	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	err := app.init(fileName, os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.do()
	if err != nil {
		if errors.Is(err, ErrNotAllowedAction) {
			fmt.Println(usage)
			return
		}
		fmt.Println(err)
		return
	}

}
