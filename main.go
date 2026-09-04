package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("todo-cli: укажите команду")
		return
	}

	switch os.Args[1] {
	case "done":
		fmt.Println("задача отмечена выполненной")
	default:
		fmt.Println("неизвестная команда")
	}
}