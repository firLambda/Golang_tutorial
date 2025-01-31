package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	logger := log.New(f, "", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("message")
	log.Println()

	_, err = os.Open("fdasfdsa")
	if err != nil {
		logger.Fatalln("!!Exit", err)
	}
}
