package utils

import (
	"fmt"
	"strings"
)

type wrap struct {
	logList []string
}

func NewWrapLogger() wrap {
	return wrap{}
}

func (w *wrap) Wrap(logMsg string) {

	w.logList = append(w.logList, logMsg)
}

func (w *wrap) Print() {
	for i, v := range w.logList {
		bar := strings.Repeat("-", i+1)
		fmt.Printf("|%s %s\n", bar, v)
	}
}

func RedLogger[T int | string](msg T) {

	newMsg := ""
	switch v := any(msg).(type) {
	case int:
		newMsg = fmt.Sprintf("%d", v)

	case string:
		newMsg = fmt.Sprintf("%s", v)

	default:
		panic("Unspported type")
	}

	red := "\033[31m"
	reset := "\033[0m"

	fmt.Printf("%sTime Duration: %s%s\n", red, newMsg, reset)
}
