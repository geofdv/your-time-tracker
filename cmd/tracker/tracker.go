package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	ErrNotAllowedAction = errors.New("not allowed action")
	ErrAlreadyTracked   = errors.New("already tracked")
	ErrNotTracked       = errors.New("not tracked yet")
)

func (app *application) do() error {
	var err error

	switch app.action {
	case "start":
		err = app.startTracking()
	case "stop":
		err = app.stopTracking()
	case "status":
		if app.isTracked() {
			fmt.Println("currently tracked")
		} else {
			fmt.Println("currently not tracked")
		}
	default:
		fmt.Println(usage)
		err = app.clean()
	}
	return err

}

func (app *application) startTracking() error {
	if app.isTracked() {
		return ErrAlreadyTracked
	}

	return app.writeRecord(NewRecord(app.action))
}

func (app *application) stopTracking() error {
	if !app.isTracked() {
		return ErrNotTracked
	}

	r, err := app.readRecord()
	if err != nil {
		return err
	}

	start, err := app.getTime(r)
	if err != nil {
		return err
	}

	d := time.Now().Sub(start)

	fmt.Fprintf(os.Stdout, "elapsed time: %s\n", d.Truncate(time.Minute))

	return app.clean()
}

func (app *application) isTracked() bool {
	r, err := app.readRecord()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return false
		} else {
			panic(err)
		}
	}

	return r.Event == "start"
}
