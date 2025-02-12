package lib

import (
	"net/url"
)

func IsURL(link string) bool {
	p, err := url.Parse(link)
	if err != nil {
		return false
	}
	if p.Scheme != "http" && p.Scheme != "https" {
		return false
	}
	return true
}
