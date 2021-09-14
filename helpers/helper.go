package helpers

import (
	"strings"
)

func UrlContain(reqUrl, url string) bool {
	return strings.Contains(reqUrl, url)
}