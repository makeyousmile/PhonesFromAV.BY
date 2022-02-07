package main

import "log"

func main() {

	for _, id := range GetIds(ScrapPage("1")) {
		log.Print(GetPhone(id))
	}

}
