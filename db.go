package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("db > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
		line := scanner.Text()
		if line == "exit" {
			fmt.Println("bye felicia")
			break
		}
		fmt.Println(line)
	}
}
