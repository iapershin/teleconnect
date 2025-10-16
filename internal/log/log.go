package log

import (
	"github.com/fatih/color"
)

func LogSuccess(message string) {
	color.New(color.FgGreen, color.Bold).Println(message)
}

func LogError(message string) {
	color.New(color.FgHiRed, color.Bold).Println(message)
}
