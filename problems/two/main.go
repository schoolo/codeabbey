package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num_elements, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	elements := scanner.Text()
	sum := 0
	for _, element := range strings.Split(elements, " ") {
		increment, _ := strconv.Atoi(element)
		sum += increment
	}
	fmt.Println(num_elements, ":", elements, ":", sum)
}
