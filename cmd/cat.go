package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Cat(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	files := req.Form["target"]
	file := files[0]
	data, err := ioutil.ReadFile(file)
	if nil != err {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprint(w, string(data))
	}
}
