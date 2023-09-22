package apicaller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/glebpepega/proj/decoder"
	"github.com/glebpepega/proj/person"
)

func CallAPI(p *person.Person, url string) (*person.Person, error) {
	resp, err := http.Get(url + p.Name)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	decoder.DecodeFromJSON(resp.Body, &p)
	if p.Count == 0 {
		return nil, fmt.Errorf("invalid name")
	} else {
		if len(p.Country) > 0 {
			p.OriginCountry = p.Country[0].Country_id
		}
		return p, nil
	}
}
