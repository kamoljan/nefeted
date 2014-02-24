package json

import (
	"encoding/json"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Msg struct {
	Status  string
	Message interface{}
}

func Message(status string, message interface{}) []byte {
	m := Msg{
		Status:  status,
		Message: message,
	}
	b, err := json.Marshal(m)
	check(err)
	return b
}
