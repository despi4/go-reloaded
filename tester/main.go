package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "12"

	b, _ := strconv.Atoi(a)

	fmt.Println(b)
}
