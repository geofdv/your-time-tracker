package main

import (
	"os"
)

func (app *application) init(fname, a string) error {
	const (
		flags = os.O_APPEND | os.O_CREATE | os.O_RDWR
		perms = 0644
	)

	f, err := os.OpenFile(fname, flags, perms)
	if err != nil {
		return err
	}

	app.action = a
	app.storage.name = fname
	app.storage.file = f

	return nil
}

func (app *application) writeRecord(r Record) error {
	err := app.writeJSON(app.storage.file, r)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) readRecord() (Record, error) {
	ret, err := app.storage.file.Seek(0, 0)
	if err != nil || ret < 0 {
		return Record{}, err
	}

	var r Record

	err = app.readJSON(app.storage.file, &r)
	if err != nil {
		return Record{}, err
	}

	return r, nil
}

func (app *application) clean() error {
	app.storage.file.Close()
	return os.Remove(app.storage.name)
}
