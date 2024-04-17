package amap

import "encoding/json"

type Error struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	InfoCode string `json:"infocode"`
}

func (error *Error) Error() string {
	return error.Info
}

func (error *Error) IsSuccess() bool {
	return error.Status == "1"
}

type String string

func (s *String) UnmarshalJSON(data []byte) error {
	switch data[0] {
	case '"':
		var ns string
		if err := json.Unmarshal(data, &ns); err != nil {
			return err
		}
		*s = String(ns)
	}
	return nil
}
