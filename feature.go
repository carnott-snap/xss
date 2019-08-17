package xss

import (
	"io/ioutil"
	"net/http"
)

func Test() {
	r, _ := http.Get("")
	ioutil.ReadAll(r.Body)
}
