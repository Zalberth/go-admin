package action

import (
	"encoding/json"
	"strings"
)

type AjaxData map[string]interface{}

func NewAjaxData() AjaxData {
	return AjaxData{"ids": "{%ids}"}
}

func (a AjaxData) Add(m map[string]interface{}) AjaxData {
	for k, v := range m {
		a[k] = v
	}
	return a
}

func (a AjaxData) JSON() string {
	b, _ := json.Marshal(a)
	s := strings.Replace(string(b), `"{%ids}"`, "{%ids}", -1)
	s = strings.Replace(s, `"{%id}"`, "{%id}", -1)
	return s
}
