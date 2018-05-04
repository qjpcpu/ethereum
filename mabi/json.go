package mabi

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

func (obj JSONObj) Get(name string) interface{} {
	return obj[name]
}

func (obj JSONObj) String() string {
	data, _ := json.Marshal(obj)
	return string(data)
}
