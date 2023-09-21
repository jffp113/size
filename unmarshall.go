package size

import (
	"encoding/json"
	"errors"
)

func (d Size) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Size) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case float64:
		*d = Size(value)
		return nil
	case string:
		tmp, err := Parse(value)
		if err != nil {
			return err
		}
		*d = tmp
		return nil
	default:
		return errors.New("invalid size")
	}
}
