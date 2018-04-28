package abi

import (
	"encoding/json"
)

type JSONObj map[string]interface{}

func NewJSONObj() JSONObj {
	return make(JSONObj)
}

func (obj JSONObj) Set(name string, val interface{}) {
	obj[name] = val
}

func (obj JSONObj) String() string {
	data, _ := json.Marshal(obj)
	return string(data)
}
