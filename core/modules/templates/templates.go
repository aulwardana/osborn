package template

import (
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func ServeAngular(r *mux.Router, path string, dir string) {
	r.PathPrefix(path).Handler(http.FileServer(http.Dir(dir)))
	go func() {
		cmd := exec.Command("tsc", "-w", "-p", "assets")
		cmd.Start()
	}()
}
