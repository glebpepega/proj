package decoder

import (
	"encoding/json"
	"io"
	"log"
)

func DecodeFromJSON(r io.Reader, obj any) {
	d := json.NewDecoder(r)
	if err := d.Decode(obj); err != nil {
		log.Println(err)
	}
}
