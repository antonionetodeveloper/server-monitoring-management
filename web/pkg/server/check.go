package server

import (
	"SMM/pkg/logs"
	"fmt"
	"net/http"
)

func CheckServers() {
	sites := GetServers()

	for _, url := range sites {
		resp, err := http.Get(url)

		if err != nil || resp.StatusCode != 200 {
			fmt.Println("[ ??? ]", url, "is uncessable.")
			logs.Register(url, 400)
		} else {
			fmt.Println("[", resp.StatusCode, "]", url, "is ok.")
			logs.Register(url, resp.StatusCode)
		}
	}
}
