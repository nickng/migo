package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nickng/migo"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	parsed, err := migo.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(parsed.String())
}
