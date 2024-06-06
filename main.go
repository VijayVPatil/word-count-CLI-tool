package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	c string
)

func init() {
	flag.StringVar(&c, "c", "", "Option c to output number of bytes in file")
	flag.Parse()
}

func main() {
	sb, err1 := os.ReadFile(c)

	if err1 != nil {
		log.Fatal(err1)
	}

	count := len(sb)

	fmt.Println("The total bytes in file are", count)
}
