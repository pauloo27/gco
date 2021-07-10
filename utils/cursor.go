package utils

import "fmt"

func MoveCursorUp(lineCount int) {
	fmt.Printf("\033[%dF", lineCount)
}

func ClearLine() {
	fmt.Print("\033[K")
}
