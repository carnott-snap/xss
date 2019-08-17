package xss

import (
	"io/ioutil"
	"net/http"
)

// Test is a function that test the GolangCI platform.
func Test() {
	r, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
}
