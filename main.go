package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix(os.Args[0], "go-build"))
}
