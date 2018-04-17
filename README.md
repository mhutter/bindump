# bindump

Reads data from `stdin`; and prints the binary representation of each byte.

## Usage

    $ echo 'Hello, World!' | bd
    01001000 01100101 01101100 01101100
    01101111 00101100 00100000 01010111
    01101111 01110010 01101100 01100100
    00100001 00001010

## Installation

    go get -u github.com/mhutter/bindump/...
