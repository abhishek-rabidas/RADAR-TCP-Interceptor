package parsing

import (
	"fmt"
	"radar/parsing/entity"
	"strconv"
)

func ParseByteStreams(streamBytes []byte) {

	i := 0

	for {
		if i+2 >= len(streamBytes) {
			break
		}

		var stream entity.Stream = entity.Stream{}

		stream.Id = streamBytes[i : i+2] //0....1
		i += 2                           //2

		stream.Length = streamBytes[i : i+1] //2
		i += 1                               //3

		length, _ := strconv.Atoi(string(stream.Length)) //8
		j := i + length                                  // 10
		stream.Payload = streamBytes[i:j]                //3....10
		i = j                                            //11

		fmt.Printf("%+v\n", stream)

	}

}
