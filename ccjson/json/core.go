package json

import (
	"errors"
	"fmt"
	"strings"
)

type Array []any
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
				return i, "", errors.New("expected \", but ran out of string")
			}
		}
		return j + 1, jsonString[i+1 : j], nil
	} else { // not a string
		return i, "", errors.New("expected \", this is not a string")
	}
}

func ParseObject(i int, jsonString string) (int, Object, error) {
	object := Object{}
	totalLength := len(jsonString)
	if jsonString[i] == '{' {
		j := i + 1
		for jsonString[j] != '}' {
			if j != i+1 {
				if jsonString[j] != ',' {
					return i, nil, errors.New(fmt.Sprintf("expected , at %d", j))
				} else {
					j += 1
				}
			}

			// check key
			k, key, err := ParseString(j, jsonString)
			if err != nil {
				return i, nil, errors.New(fmt.Sprintf("not a valid key - expected a string at %d", j))
			}
			if k >= totalLength || jsonString[k] != ':' {
				return i, nil, errors.New(fmt.Sprintf("expected : at %d", k))
			}
			// parse based on value type
			j = k + 1
			if j >= totalLength {
				return i, nil, errors.New(fmt.Sprintf("json ended without value for key at %d", j))
			}

			if jsonString[j] == '"' {
				// string
				k_, value, err := ParseString(j, jsonString)
				if err != nil {
					return i, nil, errors.New(fmt.Sprintf("not a valid value string at %d", j))
				}
				object[key] = value
				k = k_
			} else if jsonString[j] == '[' {
				// list
				k_, value, err := ParseArray(j, jsonString)
				if err != nil {
					return i, nil, errors.New(fmt.Sprintf("not a valid array at %d", j))
				}
				object[key] = value
				k = k_
			} else if jsonString[j] == '{' {
				// object
				k_, value, err := ParseObject(j, jsonString)
				if err != nil {
					return i, nil, errors.New(fmt.Sprintf("not a valid object at %d", j))
				}
				object[key] = value
				k = k_
			} else if jsonString[j] == ' ' || jsonString[j] == '\n' {
				// space or newline
				k = j + 1
			} else {
				// unknown character
				return i, nil, errors.New(fmt.Sprintf("invalid character at %d", j))
			}

			j = k
			if j == len(jsonString) {
				return i, nil, errors.New("expected }, but ran out of string")
			}
		}
		return j + 1, object, nil
	} else {
		return i, nil, errors.New("expected {, this is not an object")
	}
}

func ParseArray(i int, jsonString string) (int, Array, error) {
	array := Array{}
	totalLength := len(jsonString)
	if jsonString[i] == '[' {
		j := i + 1
		for jsonString[j] != ']' {
			if j != i+1 {
				if jsonString[j] != ',' {
					return i, nil, errors.New(fmt.Sprintf("expected , at %d", j))
				} else {
					j += 1
				}
			}

			// parse based on value type
			k := j
			if jsonString[j] == '"' {
				// string
				k_, value, err := ParseString(j, jsonString)
				if err != nil {
					return i, nil, errors.New(fmt.Sprintf("not a valid value string at %d", j))
				}
				array = append(array, value)
				k = k_
			} else if jsonString[j] == '[' {
				// list
				k_, value, err := ParseArray(j, jsonString)
				if err != nil {
					return i, nil, errors.New(fmt.Sprintf("not a valid array at %d", j))
				}
				array = append(array, value)
				k = k_
			} else if jsonString[j] == '{' {
				// object
				k_, value, err := ParseObject(j, jsonString)
				if err != nil {
					return i, nil, errors.New(fmt.Sprintf("not a valid object at %d", j))
				}
				array = append(array, value)
				k = k_
			} else if jsonString[j] == ' ' || jsonString[j] == '\n' {
				// space or newline
				k = j + 1
			} else {
				// unknown character
				return i, nil, errors.New(fmt.Sprintf("invalid character at %d", j))
			}

			j = k
			if j == totalLength {
				return i, nil, errors.New("expected ], but ran out of string")
			}
		}

		return j + 1, array, nil
	} else {
		return i, nil, errors.New("expected [, this is not an list")
	}
}
