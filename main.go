package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func compileFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := "test.go"
	_, errFileExist := os.Stat(path)
	if os.IsNotExist(errFileExist) {
		fmt.Println("file does not exist")
	} else {
		fmt.Println("file exist")
		os.Truncate("test.go", 0)
		os.Remove(path)
	}

	ioutil.WriteFile("test.go", ([]byte(r.Form.Get("formData"))), 0644)
	cmd := exec.Command("go", "run", "test.go")
	cmd.Stdin = strings.NewReader(r.Form.Get("formData"))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		w.Write([]byte(stderr.String()))
	} else {
		w.Write([]byte(out.String()))
	}
	os.Remove(path)
}

func main() {
	http.HandleFunc("/compileFile", compileFile)
	http.Handle("/", http.FileServer(http.Dir("templates")))

	log.Println("Listening....")
	http.ListenAndServe(":3002", nil)
}
