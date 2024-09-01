package utils

import (
	"testing"
)

func Test_logger(t *testing.T) {

	RedLogger(123)
	RedLogger("123")
}

func Test_wrapLogger(t *testing.T) {

	w := NewWrapLogger()

	w.Wrap("hello")
	w.Wrap("world")
	w.Wrap("leedonggyu")
	w.Wrap("togogogo")

	w.Print()
}
