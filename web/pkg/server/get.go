package server

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetServers() (sites []string) {
	file, err := os.Open("../pkg/server/servers.txt")
	if err != nil {
		fmt.Println("[pkg/server/add AddServer() os.Open('../pkg/server/servers.txt')] - Error opening file")
		fmt.Println("Error ---> ", err)
		os.Exit(-1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(line))
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("[pkg/server/add AddServer() reader.ReadString('...')] - Error reading file")
			fmt.Println("Error ---> ", err)
			os.Exit(-1)
		}
	}

	return
}

func PrintAllServers() {
	sites := GetServers()

	for _, url := range sites {
		fmt.Println(url)
	}
}
