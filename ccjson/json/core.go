package json

import (
	"errors"
	"fmt"
	"strings"
)

type List []any
type Object map[any]any

func ParseString(i int, jsonString string) (int, string, error) {
	// is a string
	if jsonString[i] == '"' {
		j := i + 1
		var stringBuilder strings.Builder
		for jsonString[j] != '"' {
			stringBuilder.WriteByte(jsonString[j])

			j += 1
			if j == len(jsonString) {
				return -1, "", errors.New("expected \", but ran out of string")
			}
		}
		return j + 1, jsonString[i+1 : j], nil
	} else { // not a string
		return i, "", errors.New("expected \", this is not a string")
	}
}

func ParseObject(i int, jsonString string) (Object, error) {
	object := Object{}
	totalLength := len(jsonString)
	if jsonString[i] == '{' {
		j := i + 1
		for jsonString[j] != '}' {
			if j != i+1 {
				if jsonString[j] != ',' {
					return nil, errors.New(fmt.Sprintf("expected , at %d", j))
				} else {
					j += 1
				}
			}

			// check key
			k, key, err := ParseString(j, jsonString)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("not a valid key - expected a string at %d", j))
			}
			if k >= totalLength || jsonString[k] != ':' {
				return nil, errors.New(fmt.Sprintf("expected : at %d", k))
			}
			// parse based on value type
			j = k + 1
			if j >= totalLength {
				return nil, errors.New(fmt.Sprintf("json ended without value for key at %d", j))
			}

			// string
			if jsonString[j] == '"' {
				k_, value, err := ParseString(j, jsonString)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("not a valid value string at %d", j))
				}
				object[key] = value
				k = k_
			}
			j = k
			if j == len(jsonString) {
				return nil, errors.New("expected }, but ran out of string")
			}
		}
		return object, nil
	} else {
		return nil, errors.New("expected {, this is not an object")
	}
}
