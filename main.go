package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	Port      string
	Hooks     []hook
}

func (c configObject) command(url string) {
	for _, hook := range c.Hooks {
		if hook.Full == url {
			cmdlet := strings.Fields(hook.Command)
			if len(cmdlet) == 1 {
				cmd := exec.Command(cmdlet[0])
				cmd.Run()
				log.Println("Ran: " + hook.Command)
			} else {
				cmd := exec.Command(cmdlet[0], cmdlet[1:]...)
				cmd.Run()
				log.Println("Ran: " + hook.Command)
			}
		}
	}
}

func (h hook) print() {
	log.Println("Configured new hook!")
	log.Println("\tName:" + h.Name)
	log.Println("\tCommand:" + h.Command)
	log.Println("\tURL:" + h.Full)
}

func main() {

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	f, err := os.OpenFile("./webhooks.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("\n\n\n\t")

	var config configObject

	json.Unmarshal(file, &config)

	hooks := config.Hooks

	for k := range hooks {
		hp := &hooks[k]
		hp.Full = "/" + config.Secretkey + "/" + hooks[k].URL

		http.HandleFunc(hooks[k].Full, func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("Recieved POST response from: " + r.RemoteAddr)
					log.Println("Recieved POST response: " + string(body))
				}
			}

			config.command(r.RequestURI)
		})
		hooks[k].print()
	}
	port := ":" + string(config.Port)
	http.ListenAndServe(port, nil)

}
