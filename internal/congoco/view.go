package congoco

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ViewType string

const (
	TXT  ViewType = "txt"
	INI  ViewType = "ini"
	JSON ViewType = "json"
)

type Output map[string]any

type View struct {
	Show func(output Output)
}

func NewView(vType ViewType) (*View, error) {
	v := View{}
	vt := ViewType(vType)
	switch vt {
	case TXT:
		v.Show = showTxt
	case INI:
		v.Show = showIni
	case JSON:
		v.Show = showJson
	default:
		return nil, fmt.Errorf("Unknown formatter type: %s", vType)
	}
	return &v, nil
}

func showTxt(output Output) {
	for key, val := range output {
		fmt.Printf("%s: %s\n", key, val)
	}
}

func showIni(output Output) {
	for key, val := range output {
		key = strings.ToUpper(key)
		key = strings.ReplaceAll(key, " ", "_")
		fmt.Printf("%s=\"%s\"\n", key, val)
	}
}

func showJson(output Output) {
	jsonOutput := Output{}
	for key, val := range output {
		jsonKey := strings.ToLower(key)
		jsonKey = strings.ReplaceAll(jsonKey, " ", "_")
		jsonOutput[jsonKey] = val
	}
	b, err := json.Marshal(jsonOutput)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
