package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	left, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	right, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(left + right)
}
