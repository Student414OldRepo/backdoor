package cmd

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "你好")
}
