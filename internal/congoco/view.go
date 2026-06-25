package congoco

import (
	"encoding/json"
	"fmt"
	"strings"

	"congoco/internal/flags"
)

type Output map[string]any

type View struct {
	Show func(output Output)
}

func NewView(fType flags.FormatterType) (*View, error) {
	v := View{}
	ft := flags.FormatterType(fType)
	switch ft {
	case flags.TXT:
		v.Show = showTxt
	case flags.INI:
		v.Show = showIni
	case flags.JSON:
		v.Show = showJson
	default:
		return nil, fmt.Errorf("Unknown formatter type: %s", fType)
	}
	return &v, nil
}

func showTxt(output Output) {
	for key, val := range output {
		switch v := val.(type) {
		case string:
			fmt.Printf("%s: %s\n", key, val)

		case []string:
			for _, s := range v {
				fmt.Printf("%s\n", s)
			}

		case map[string]map[string]string:
			fmt.Printf("%s:\n", key)
			for mk, mv := range v {
				fmt.Printf("\t%s\n", mk)
				for mmk, mmv := range mv {
					fmt.Printf("\t\t%s: %s\n", mmk, mmv)
				}
			}

		default:
			fmt.Printf("%s: %v\n", key, val)
		}
	}
}

func showIni(output Output) {
	for key, val := range output {
		key = strings.ToUpper(key)
		key = strings.ReplaceAll(key, " ", "_")
		switch v := val.(type) {
		case string:
			fmt.Printf("%s = %s\n", key, val)

		case []string:
			fmt.Printf("%s = %s\n", key, strings.Join(v, ";"))

		case map[string]map[string]string:
			for mk, mv := range v {
				for mmk, mmv := range mv {
					fmt.Printf("%s_%s_%s = %s\n", key, strings.ToUpper(mk), strings.ToUpper(mmk), mmv)
				}
			}

		default:
			fmt.Printf("%s = %v\n", key, val)

		}
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
