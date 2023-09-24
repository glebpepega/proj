package decoder

import (
	"encoding/json"
	"io"
)

func DecodeFromJSON(r io.Reader, obj any) error {
	d := json.NewDecoder(r)
	if err := d.Decode(obj); err != nil {
		return err
	}
	return nil
}
