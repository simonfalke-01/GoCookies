package main

import "time"

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
