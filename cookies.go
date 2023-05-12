package main

import (
	_ "github.com/simonfalke-01/cbr-cli/kooky/browser/all"
	"github.com/simonfalke-01/gocookies/kooky"
	"time"
)

type jsonCookie struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	Path       string    `json:"path"`
	Domain     string    `json:"domain"`
	Expires    time.Time `json:"expires"`
	RawExpires string    `json:"rawExpires"`
	MaxAge     int       `json:"maxAge"`
	Secure     bool      `json:"secure"`
	HttpOnly   bool      `json:"httpOnly"`
	SameSite   int       `json:"sameSite"`
	Raw        string    `json:"raw"`
	Unparsed   []string  `json:"unparsed"`
	Creation   time.Time `json:"creation"`
	Container  string    `json:"container"`
}

func getAllCookies() *[]jsonCookie {
	cookie := kooky.ReadCookies(kooky.Valid)

	var jsonCookies []jsonCookie
	for _, c := range cookie {
		jsonCookies = append(jsonCookies, jsonCookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires,
			RawExpires: c.RawExpires,
			MaxAge:     c.MaxAge,
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   int(c.SameSite),
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
			Creation:   c.Creation,
			Container:  c.Container,
		})
	}

	return &jsonCookies
}
