package decorder

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"sync"
)

type User struct {
	Name string
	Age  int
}

var inputPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Reader)
	},
}

func Get(b []byte) *bytes.Reader {
	input := inputPool.Get().(*bytes.Reader)
	input.Reset(b)
	return input
}

func Put(r *bytes.Reader) {
	inputPool.Put(r)
}

func Unmarshal(b []byte, v interface{}) error {
	b2 := Get(b)
	dec := json.NewDecoder(b2)
	token, err := dec.Token()
	if token != json.Delim(123) {
		return errors.New("invalid character found, expecting { for JSON object")
	}
	if err != nil {
		return err
	}

	if u, ok := v.(*User); ok {
		u.Name = ""
		u.Age = 0

		for dec.More() {
			token, err = dec.Token()
			if err != nil {
				return err
			}

			switch token {
			case "Name":
				err = dec.Decode(&u.Name)
				if err != nil {
					return err
				}
			case "Age":
				err = dec.Decode(&u.Age)
				if err != nil {
					return err
				}
			}
		}
	}

	token, err = dec.Token()
	if token != json.Delim(125) {
		return errors.New("invalid character found, expecting } for JSON object")
	}

	Put(b2)

	return err
}

func U2(jsonData []byte) (result User) {
	start := 0
	end := 0

	for {
		for start < len(jsonData) && jsonData[start] != '{' {
			start++
		}

		for end < len(jsonData) && jsonData[end] != '}' {
			end++
		}

		nameStart := start + 1
	LOOP:
		for ; start < end; start++ {
			if jsonData[nameStart] == ',' || jsonData[nameStart] == ' ' {
				nameStart++
			}
			if jsonData[start] == ':' {
				name := string(jsonData[nameStart:start])
				if name == "\"Name\"" {
					start++
					for ; start < end; start++ {
						if jsonData[start] == '"' {
							start++
							for ; start < end; start++ {
								if jsonData[start] == '"' {
									start++
									result.Name = string(jsonData[nameStart:start])
									nameStart = start
									goto LOOP
								}
							}
						}
					}
				} else if name == "\"Age\"" {
					for digitStart := start + 1; digitStart < end; digitStart++ {
						if jsonData[digitStart] >= '0' && jsonData[digitStart] <= '9' {
							age, _ := strconv.Atoi(string(jsonData[start+1 : end]))
							result.Age = age
							break
						}
					}
				}
			}
		}

		if end == len(jsonData) {
			break
		}
		start, end = end+1, end+1
	}
	return
}
