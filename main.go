package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type hook struct {
	Name    string
	URL     string
	Command string
	Full    string
}

type configObject struct {
	Secretkey string
	Hooks     []hook
}

func (c configObject) command(url string) {
	for _, hook := range c.Hooks {
		if hook.Full == url {
			//TODO make this actually run the command
			fmt.Println(hook.Command)
		}
	}
}

func (h *hook) setFull(full string) {
	h.Full = full
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
	http.ListenAndServe(":8080", nil)

}
