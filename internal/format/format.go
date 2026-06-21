package format

import (
	"fmt"
	"strings"
)

type FormatterType string

const (
	TXT  FormatterType = "txt"
	INI  FormatterType = "ini"
	JSON FormatterType = "json"
)

type Output map[string]string

type Formatter struct {
	Render func(text Output)
}

func New(formatterType string) (*Formatter, error) {
	formatter := Formatter{}

	ft := FormatterType(formatterType)
	switch ft {
	case TXT:
		formatter.Render = renderTXT
	case INI:
		formatter.Render = renderINI
	case JSON:
		formatter.Render = renderJSON
	default:
		return nil, fmt.Errorf("Unknown formatter type: %s", formatterType)
	}

	return &formatter, nil
}

func renderTXT(output Output) {
	for key, val := range output {
		fmt.Printf("%s: %s\n", key, val)
	}
}

func renderINI(output Output) {
	for key, val := range output {
		key = strings.ToUpper(key)
		key = strings.ReplaceAll(key, " ", "_")
		fmt.Printf("%s=\"%s\"\n", key, val)
	}
}
func renderJSON(output Output) {}
