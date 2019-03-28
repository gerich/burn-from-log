package main

import (
	"flag"
	"fmt"
)

const hostRegexp = ""

func main() {
	host := flag.String("host", "", "Host for burn")
	flag.Parse()
	if *host == "" {
		panic(fmt.Errorf("Empty host"))
	}

	fmt.Println(*host)
}
