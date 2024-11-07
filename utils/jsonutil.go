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

func PrettyPrint(v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bytes))
}
