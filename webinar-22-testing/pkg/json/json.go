package json

import "encoding/json"

var jsonTest = `{"a": "b", "c": 3}`

func UnmarshalToMap() error {
	m := make(map[string]any)

	err := json.Unmarshal([]byte(jsonTest), &m)
	if err != nil {
		return err
	}

	return nil
}

type m struct {
	A string `json:"a"`
	C int    `json:"c"`
}

func UnmarshalToStruct() error {
	var m m

	err := json.Unmarshal([]byte(jsonTest), &m)
	if err != nil {
		return err
	}

	return nil
}
