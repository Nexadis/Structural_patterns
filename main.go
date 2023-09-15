package main

import (
	"fmt"
	"strings"
)

func main() {
	border("Decorator")
}

func border(name string) {
	line := strings.Repeat("=", 80)
	out := fmt.Sprintf("%s\n\t\t\t\t%s\n%s", line, name, line)
	fmt.Println(out)
}
