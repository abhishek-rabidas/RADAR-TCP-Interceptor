package parsing

import (
	"fmt"
	"golang.org/x/text/unicode/bidi"
	"radar/parsing/entity"
	"strconv"
)

func ParseStreams(streamBytes string) []entity.Stream {

	i := 0

	var output []entity.Stream

	for {

		if i+4 >= len(streamBytes) {
			break
		}

		var stream entity.Stream = entity.Stream{}

		stream.Id = streamBytes[i : i+4]
		i += 4

		stream.Length, _ = strconv.Atoi(streamBytes[i : i+2])
		i += 2

		j := i + (stream.Length * 2)
		stream.Payload = bidi.ReverseString(streamBytes[i:j]) //reversing the payload to make it big endian
		i = j

		fmt.Printf("%+v\n", stream)

		output = append(output, stream)

	}

	return output
}
