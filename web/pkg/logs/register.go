package logs

import (
	"fmt"
	"os"
	"time"
)

func Register(site string, status any) {
	file, err := os.OpenFile("../pkg/logs/logs.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("[pkg/logs/register Register() os.Open('../pkg/logs/logs.txt')] - Error opening file")
		fmt.Println("Error ---> ", err)
		os.Exit(-1)
	}
	defer file.Close()

	if status == 200 {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " -> [ 200 ] " + site + " is ok!\n")
	} else {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " -> [ ??? ] " + site + " IS DOWN\n")
	}

}
