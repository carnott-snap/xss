package xss

import (
	"io/ioutil"
	"net/http"
)

func Test() {
	r, _ := http.Get("https://example.com")
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
}
