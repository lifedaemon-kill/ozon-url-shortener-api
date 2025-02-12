package lib

import (
	"net/url"
)

func IsURL(link string) bool {
	_, err := url.Parse(link)
	if err != nil {
		return false
	}
	return true
}
