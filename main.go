package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type hook struct {
	Name    string
	URL     string
	Command string
	Full    string
}

type configObject struct {
	Secretkey string
	Port      int
	Hooks     []hook
}

func (c configObject) command(url string) {
	for _, hook := range c.Hooks {
		if hook.Full == url {
			cmdlet := strings.Fields(hook.Command)
			if len(cmdlet) == 1 {
				cmd := exec.Command(cmdlet[0])
				cmd.Run()
			} else {
				cmd := exec.Command(cmdlet[0], cmdlet[1:]...)
				cmd.Run()
			}
		}
	}
}

func main() {

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var config configObject

	json.Unmarshal(file, &config)

	hooks := config.Hooks

	for k := range hooks {
		hp := &hooks[k]
		hp.Full = "/" + config.Secretkey + "/" + hooks[k].URL

		http.HandleFunc(hooks[k].Full, func(w http.ResponseWriter, r *http.Request) {
			config.command(r.RequestURI)
		})
		fmt.Println("Configured new hook!")
		fmt.Println("\tName:\t", hooks[k].Name)
		fmt.Println("\tURL:\t", hooks[k].Full)
		fmt.Println("\tCommand:", hooks[k].Command)
	}
	port := ":" + string(config.Port)
	http.ListenAndServe(port, nil)

}
