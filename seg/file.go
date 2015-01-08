package seg

import (
	"io/ioutil"
)

func ReadAll(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(b)
}
