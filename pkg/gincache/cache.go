package gincache

import "encoding/json"

type Cache struct {
	Status int    `json:"status"`
	Body   []byte `json:"body"`
}

func (c Cache) ToBytes() ([]byte, error) {
	return json.Marshal(c)
}
