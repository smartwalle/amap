package amap

const (
	kAPIDomain = "https://restapi.amap.com/v3"
)

type Error struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	InfoCode string `json:"infocode"`
}

func (this *Error) Error() string {
	return this.Info
}

func (this *Error) IsSuccess() bool {
	return this.Status == "1"
}

type String string

func (this String) MarshalJSON() ([]byte, error) {
	return []byte(this), nil
}

func (this *String) UnmarshalJSON(data []byte) error {
	var l = len(data)
	if data[0] == '[' && data[l-1] == ']' {
		return nil
	}
	if data[0] == '"' {
		*this = String(data[1 : l-1])
	}
	return nil
}
