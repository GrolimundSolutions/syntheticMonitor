package util

import "fmt"

// URLBuilder returns a Url buildet with the given Params
func URLBuilder(host string, port uint16, path string) (string, error) {
	fullURL := fmt.Sprintf("%s:%d%s", host, port, path)
	return fullURL, nil
}
