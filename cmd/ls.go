package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type FileInfo struct {
	Name string
	Size int64
	Mode os.FileMode
	Time time.Time
	Dir  bool
}

func Ls(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	targets := req.Form["target"]
	target := targets[0]
	fmt.Println(target)
	files, err := ioutil.ReadDir(target)
	datas := make([]*FileInfo, 0)
	for _, file := range files {
		fileinfo := &FileInfo{
			file.Name(),
			file.Size(),
			file.Mode(),
			file.ModTime(),
			file.IsDir(),
		}
		datas = append(datas, fileinfo)
	}
	// fmt.Println(datas)

	if nil != err {
		fmt.Fprint(w, err)
	} else {
		data, err := json.Marshal(datas)
		if nil != err {
			fmt.Fprint(w, err)
		} else {
			fmt.Fprint(w, string(data))
		}
	}
}
