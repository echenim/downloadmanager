package main

import "log"

func main() {
	s := "error error error"
	log.Printf("%v\n", s)
	log.Fatalf("error occured while retriving the file: %s\n\n", s)
	log.Printf("%v\n", s)
}
