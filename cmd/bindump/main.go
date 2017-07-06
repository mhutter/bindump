package main

import (
	"os"

	"github.com/mhutter/bindump"
)

func main() {
	bindump.Dump(os.Stdin, os.Stdout)
}
