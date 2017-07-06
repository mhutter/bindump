package bindump

import (
	"bufio"
	"fmt"
	"io"
)

// Dump reads data from `in` and prints a binary representation to `out`.
func Dump(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanBytes)

	i := 1

	for scanner.Scan() {
		fmt.Fprintf(out, "%08b ", scanner.Bytes()[0])
		if i%4 == 0 {
			fmt.Fprint(out, "\n")
		}

		i++
	}
	fmt.Fprint(out, "\n")
}
