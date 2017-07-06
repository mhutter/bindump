package bindump

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	input := strings.NewReader("Hello, World!")
	var output bytes.Buffer

	expected := "01001000 01100101 01101100 01101100 \n" +
		"01101111 00101100 00100000 01010111 \n" +
		"01101111 01110010 01101100 01100100 \n" +
		"00100001 \n"

	Dump(input, &output)

	if output.String() != expected {
		t.Errorf(
			"got %#v\n     but expected %#v\n",
			output.String(),
			expected)
	}
}
