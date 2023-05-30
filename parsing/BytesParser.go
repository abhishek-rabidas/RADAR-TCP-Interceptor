package parsing

import (
	"fmt"
	"radar/parsing/entity"
	"strconv"
)

func ParseByteStreams(streamBytes []byte) {

	i := 0

	//fmt.Printf("%x\n", streamBytes)

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

		id := fmt.Sprintf("%x", stream.Id)

		if id == "0502" {
			//fmt.Printf("%d\n", len(stream.Payload))
			//fmt.Printf("1: {%08b}: ", stream.Payload)
			//ParseBinaryObjectData(stream.Payload)
			//fmt.Printf("2: {%08b}: ", stream.Payload)
			ParseBinaryObjectData2(stream.Payload)
			//fmt.Println()
		}

	}

}

func parseStatusMessage(payload []byte) {
	//a := binary.LittleEndian.Uint32(payload[4:8])

	fmt.Printf("\nTime: %08b\n", payload[4:8])
}

func reverse(array []byte) []byte {

	for i := 0; i < len(array)/2; i++ {
		temp := array[i]
		array[i] = array[len(array)-1-i]
		array[len(array)-1-i] = temp
	}

	return array
}

func ParseBinaryObjectData2(payload []byte) {
	stream := ""

	for _, by := range reverse(payload) {
		stream += fmt.Sprintf("%08b", by)
	}

	var object entity.ObjectData = entity.ObjectData{}

	//x_coordinate
	result, err := strconv.ParseUint(stream[1:14], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.X = float64(result) - 4096

	//y_coordinate
	result, err = strconv.ParseUint(stream[14:27], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Y = float64(result) - 4096

	//X_speed
	result, err = strconv.ParseUint(stream[27:38], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.XSpeed = (float64(result) * 0.1) - 1024

	//Y_speed
	result, err = strconv.ParseUint(stream[38:49], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.YSpeed = (float64(result) * 0.1) - 1024

	//Length
	result, err = strconv.ParseUint(stream[49:56], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Length = (float64(result) * 0.2) - 0

	//ObjectId
	result, err = strconv.ParseUint(stream[56:64], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.ObjectId = result

	fmt.Printf("%+v\n", object)

	//stream = strings.ReplaceAll(stream, " ", "")

}
