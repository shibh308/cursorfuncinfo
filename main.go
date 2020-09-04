package main

import (
	"fmt"
	"github.com/shibh308/cursorfuncinfo/run"
	"os"
)

func main() {
	str, err := run.Run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(str)
	os.Exit(0)
}
