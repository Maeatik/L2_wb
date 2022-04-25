package main

import "fmt"

type book struct {
	title       string
	author      string
	writingDone bool
	publishDone bool
	salesDone   bool
	readingDone bool
}

func (b book) string() string {
	return fmt.Sprintf("%b", b)
}
