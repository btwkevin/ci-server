package main

import (
	"fmt"
	"net/http"

	"github.com/btwkevin/ci-server/handlers"
)

func main() {
	http.HandleFunc("/webhook", handlers.WebHook)
	fmt.Println("CI Server Listen : 8080")
	http.ListenAndServe(":8080", nil)
}
