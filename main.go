package main

import (
	"fmt"
	"log"

	"github.com/aristorinjuang/reverse"
)

func main() {
	hello := []byte("Hello World!")
	list, err := reverse.Init([]byte("Hello World!"))

	if err != nil {
		log.Fatal(err)
	}

	reversedHello, err := reverse.String(hello)

	if err != nil {
		log.Fatal(err)
	}

	reversedList, err := reverse.LinkedList(list)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(reversedHello))
	fmt.Println(reverse.Print(reversedList))
}
