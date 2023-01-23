/*

Implement run-length encoding and decoding.

Run-length encoding (RLE) is a simple form of data compression, where runs (consecutive data elements) are replaced by just one data value and count.

For example we can represent the original 53 characters with only 13.

"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"  ->  "12WB12W3B24WB"

RLE allows the original data to be perfectly reconstructed from the compressed data, which makes it a lossless data compression.

"AABCCCDEEEE"  ->  "2AB3CD4E"  ->  "AABCCCDEEEE"

For simplicity, you can assume that the unencoded string will only contain the letters A through Z (either lower or upper case) and whitespace. This way data to be encoded will never contain any numbers and numbers inside data to be decoded always represent the count for the following character.

*/

package main

import (
	"fmt"

	"strconv"

	"strings"
)

func RLE(s string) (e string) {

	for len(s) > 0 {

		l := s[0]

		slen := len(s)

		s = strings.TrimLeft(s, string(l))

		if a := slen - len(s); a > 1 {

			e += strconv.Itoa(a)

		}

		e += string(l)

	}

	return e

}

func main() {

	fmt.Println(RLE("WERGHH"))

}
