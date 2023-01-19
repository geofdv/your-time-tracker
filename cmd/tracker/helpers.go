package main

import (
	"encoding/json"
	"io"
	"time"
)

func (app *application) writeJSON(w io.Writer, data any) error {
	var line []byte

	line, err := json.Marshal(data)
	if err != nil {
		return err
	}

	n, err := w.Write(append(line, '\n'))
	if err != nil || n <= 0 {
		return err
	}

	return nil
}

func (app *application) readJSON(r io.Reader, dst any) error {
	dec := json.NewDecoder(r)

	err := dec.Decode(&dst)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) getTime(r Record) (time.Time, error) {
	return time.Parse(time.RFC822, r.CreatedAt)
}
