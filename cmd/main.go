package main

import (
	"flag"
	"fmt"

	"github.com/k3forx/goqa/slice"
)

func main() {
	var s = flag.String("q", "", "quiz name")
	flag.Parse()

	switch *s {
	case "slice":
		slice.Run()
	default:
		fmt.Printf("need to specify flag")
	}
}
