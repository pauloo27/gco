package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/c-bata/go-prompt"
)

func PrettyPrint(out prompt.ConsoleWriter, components ...interface{}) {
	for _, component := range components {
		switch component := component.(type) {
		case string:
			out.WriteStr(component)
		case int:
			out.WriteStr(strconv.Itoa(component))
		case prompt.Color:
			c := component
			bold := true
			if c == prompt.DefaultColor {
				bold = false
			}
			out.SetColor(component, prompt.DefaultColor, bold)
		}
	}
	err := out.Flush()
	if err != nil {
		fmt.Println("Something went wrong while writing to output console")
		os.Exit(-1)
	}
}
