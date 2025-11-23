package handlers

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func WebHook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))

	go func() {
		cmd := exec.Command("../deploy.sh")

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			log.Println("deploy failed: ", err)
		} else {
			log.Println("deploy successful")
		}
	}()
}
