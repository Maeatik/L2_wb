package main

import (
	"os"
	"testing"
)

func TestAfterFixed(t *testing.T) {
	fg := &flags{
		After: 2,
		Fixed: true,
	}
	pattern := "lgnu9d"
	expect := "+ lgnu9d\n lignum\n magNum"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestBeforeIgnoreCase(t *testing.T) {
	fg := &flags{
		Before:     2,
		IgnoreCase: true,
	}
	pattern := "lgnu9d"
	expect := " gnu\n interregnum\n+ lgnu9d"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}

}

func TestCount(t *testing.T) {
	fg := &flags{Count: true}
	pattern := "gnu"
	expect := "7 совпадений"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestLineNumberContext(t *testing.T) {
	fg := &flags{
		Context: 1,
		LineNum: true,
	}
	pattern := "l"
	expect := "3  interregnum\n4 + lgnu9d\n5 + lignum\n6  magNum"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestInvertIgnoreCase(t *testing.T) {
	fg := &flags{
		Invert:     true,
		IgnoreCase: true,
	}
	pattern := "in"
	expect := "cygnus\ngnu\nlgnu9d\nlignum\nmagNum\nmagnuson\nsphagnum"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}
