package main

import (
	"flag"
	"fmt"
	"gocloud/cmd"
	"net/http"
)

var serverPort *int

func main() {
	serverPort = flag.Int("p", 7070, "server port")
	flag.Parse()
	fmt.Println("Welcome to GoCloud !")
	// test mysql : ok
	// mysql.TestMySQL()

	// bind handler
	http.HandleFunc("/login", cmd.Login)
	http.HandleFunc("/ls", cmd.Ls)
	http.HandleFunc("/cat", cmd.Cat)
	http.HandleFunc("/exec", cmd.Exec)
	// http.HandleFunc("/login", handler.Login)
	// http.HandleFunc("/login", handler.Login)

	// start http server
	address := fmt.Sprintf("0.0.0.0:%d", *serverPort)
	err := http.ListenAndServe(address, nil)
	if nil != err {
		fmt.Println(err)
	}

}
