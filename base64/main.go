// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Encoder(str string) string {
	reader := strings.NewReader(str)
	buffer := make([]byte, 3)
	readBytes, err := reader.Read(buffer)
	var stream string
	var sixBit int64
	var table string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	for {
		if readBytes == 0 && err == io.EOF {
			break
		} else if readBytes < 3 {
			if readBytes == 1 {
				buffer[1] = 00000000
				buffer[2] = 00000000
			} else {
				buffer[2] = 0
			}
		}
		stream = fmt.Sprintf("%08b%08b%08b", buffer[0], buffer[1], buffer[2])

		for i := 0; i < 4; i++ {
			sixBit, _ = strconv.ParseInt(stream[i*6:(i*6)+6], 2, 8)
			if sixBit == 0 {
				fmt.Print("=")
			} else {
				fmt.Printf("%c", table[sixBit])
			}

		}

		readBytes, err = reader.Read(buffer)
	}
	return ""
}

func main() {
	name := "username:password"
	Encoder(name)
}
