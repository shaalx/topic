package seg

import (
	"io/ioutil"
	"net/http"
)

func ReadAll(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(b)
}

func ReadHttp(url string) string {
	resp, err := http.Get(url)
	if nil != err {
		return ""
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return ""
	}
	return string(bs)
}
