package main

import (
	"fmt"

	"log"
	"os"
)

func main() {

	file, err := os.Create("/tmp/alilog.txt")
	if err != nil {
		log.Fatal("Error while trying to create file", err)
	}

	defer file.Close()

	for i := 1; i < 10001; i++ {
		fmt.Fprintf(file, "Hi, this is line number %v\n", i)
	}

}
