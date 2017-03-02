package template

import (
	"net/http"
	"os/exec"
)

func ServeAngular() {
	http.Handle("/", http.FileServer(http.Dir("./assets/")))
	go func() {
		cmd := exec.Command("tsc", "-w", "-p", "assets")
		cmd.Start()
	}()
}
