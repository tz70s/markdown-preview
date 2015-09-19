package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func links(w http.ResponseWriter, r *http.Request) {
	out := parse()

	r.ParseForm()
	fmt.Println("method: ", r.Method)
	t, _ := template.ParseFiles(out)
	t.Execute(w, nil)
}

func open(url string) {
	var command string
	if runtime.GOOS == "darwin" {
		command = "open"
	} else if runtime.GOOS == "linux" {
		command = "xdg-open"
	}

	cmd := exec.Command(command, url)
	cmd.Run()
}

func server(port string) {
	http.HandleFunc("/", links)
	fmt.Println("Press Enter to open the web page on browser")
	url := "http://localhost" + port + "/"
	fmt.Println(url)
	fmt.Scanln()
	open(url)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
