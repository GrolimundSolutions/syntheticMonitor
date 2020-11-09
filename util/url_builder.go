package util

import "fmt"

func UrlBuilder(host string, port uint16, path string) (string, error) {
	fullUrl := fmt.Sprintf("%s:%d%s", host, port, path)
	return fullUrl, nil
}
