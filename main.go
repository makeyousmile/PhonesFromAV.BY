package main

import (
	"log"
	"strconv"
)

func main() {
	run(999)
	//getToken()
	//sendSms("test", []string{})

}
func run(pages int) {
	for i := 0; i < pages; i++ {
		for _, id := range GetIds(strconv.FormatInt(int64(i), 10)) {
			phone := GetPhone(id)
			log.Print(phone)
			writeTofile(phone + "\n")
		}
	}

}
