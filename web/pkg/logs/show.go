package logs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ShowLogs() {
	file, err := os.Open("../pkg/logs/logs.txt")
	if err != nil {
		fmt.Println("[pkg/logs/show ShowLogs() os.Open('../pkg/logs/logs.txt')] - Error opening file")
		fmt.Println("Error ---> ", err)
		os.Exit(-1)
	}
	defer file.Close()

	fmt.Println("\n-=-=-=-=-=-=-=-=-=Logs-=-=-=-=-=-=-=-=-=")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(strings.TrimSpace(line))
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("File are not valid.")
			os.Exit(-1)
		}
	}
	fmt.Println("\n-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}
