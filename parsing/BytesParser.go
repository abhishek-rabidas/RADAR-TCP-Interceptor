package parsing

import (
	"fmt"
	"radar/parsing/entity"
)

func ParseByteStreams(streamBytes []byte) {

	i := 0

	fmt.Printf("%x\n", streamBytes)

	for {
		if i+2 >= len(streamBytes) {
			break
		}

		var stream entity.Stream = entity.Stream{}

		stream.Id = streamBytes[i : i+2] //0....1
		i += 2                           //2

		stream.Length = streamBytes[i : i+1]
		length := int(streamBytes[i])
		i += 1

		j := i + length
		stream.Payload = streamBytes[i:j]
		i = j

		fmt.Printf("%+v\n", stream)

		id := fmt.Sprintf("%x", stream.Id)

		if id == "0502" {
			ParseBinaryObjectData(stream.Payload)
		}

	}

}

func ParseBinaryObjectData(payload []byte) {

}
