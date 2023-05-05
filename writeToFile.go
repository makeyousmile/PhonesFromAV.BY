package main

import (
	"fmt"
	"os"
)

func writeTofile(data string) {

	file, err := os.OpenFile("phones.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(data)

	fmt.Println("Done.")
}
