package main

import (
	"flag"
	"fmt"
)

var logger = fmt.Println

func main() {
	logger(`===== COMMANDLINE FLAG =====`)

	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("num", 731, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	logger("word", *wordPtr)
	logger("num", *numPtr)
	logger("fork", *boolPtr)
	logger("svar", svar)
	logger("tail", flag.Args())
}
