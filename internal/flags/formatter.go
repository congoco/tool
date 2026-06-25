package flags

import "fmt"

type FormatterType string

const (
	INI  FormatterType = "ini"
	JSON FormatterType = "json"
	TXT  FormatterType = "txt"
)

func (b *FormatterType) String() string {
	return string(*b)
}

func (b *FormatterType) Type() string {
	return "belong"
}

func (b *FormatterType) Set(s string) error {
	switch FormatterType(s) {
	case INI, JSON, TXT:
		*b = FormatterType(s)
		return nil
	default:
		return fmt.Errorf("invalid formatter %q (allowed: ini, json, txt)", s)
	}
}
