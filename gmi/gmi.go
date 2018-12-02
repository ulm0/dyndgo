// Get my IP

package gmi

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetIP() (string, error) {
	resp, err := http.Get("http://icanhazip.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// Get rid of newline
	ip := strings.TrimSuffix(string(respBytes), "\n")
	return ip, nil
}
