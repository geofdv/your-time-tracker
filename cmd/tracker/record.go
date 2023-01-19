package main

import (
	"time"
)

type Record struct {
	Event     string `json:"event"`
	CreatedAt string `json:"created_at"`
}

func NewRecord(event string) Record {
	return Record{
		Event:     event,
		CreatedAt: time.Now().Format(time.RFC822),
	}
}
