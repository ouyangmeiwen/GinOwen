package utils

import (
	"encoding/json"
	"fmt"
)

func ToJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

func ToRawMessage(v interface{}) (json.RawMessage, error) {
	bodyBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(bodyBytes), err
}

func StringToRawMessage(s string) json.RawMessage {
	return json.RawMessage(s)
}

func PrettyPrint(v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bytes))
}
