package parsing

import (
	"encoding/hex"
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
			fmt.Printf("1: {%08b}: ", stream.Payload)
			ParseBinaryObjectData(stream.Payload)
			fmt.Printf("2: {%08b}: ", stream.Payload)
			ParseBinaryObjectData2(stream.Payload)
			fmt.Println()
		}

	}

}

func ParseBinaryObjectData2(payload []byte) {
	stream := ""

	for _, by := range payload {
		stream += fmt.Sprintf("%08b", by)
	}

	var object entity.ObjectData = entity.ObjectData{}

	//x_coordinate
	result, err := strconv.ParseUint(stream[50:63], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.X = result

	//y_coordinate
	result, err = strconv.ParseUint(stream[37:50], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Y = result

	//X_speed
	result, err = strconv.ParseUint(stream[26:37], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.XSpeed = result

	//Y_speed
	result, err = strconv.ParseUint(stream[15:26], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.YSpeed = result

	//Length
	result, err = strconv.ParseUint(stream[8:15], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Length = result

	//ObjectId
	result, err = strconv.ParseUint(stream[0:8], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.ObjectId = result

	fmt.Printf("%+v\n", object)

	//stream = strings.ReplaceAll(stream, " ", "")

}

func ParseBinaryObjectData(payload []byte) {
	bin := hex.EncodeToString(payload)

	convertedBinary := convertToBinary(convertFormat(bin))
	//fmt.Printf("%s\n", convertedBinary)

	var object entity.ObjectData = entity.ObjectData{}

	//x_coordinate
	result, err := strconv.ParseUint(convertedBinary[1:14], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.X = result

	//y_coordinate
	result, err = strconv.ParseUint(convertedBinary[14:27], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Y = result

	//X_speed
	result, err = strconv.ParseUint(convertedBinary[27:38], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.XSpeed = result

	//Y_speed
	result, err = strconv.ParseUint(convertedBinary[38:49], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.YSpeed = result

	//Length
	result, err = strconv.ParseUint(convertedBinary[49:56], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Length = result

	//ObjectId
	result, err = strconv.ParseUint(convertedBinary[56:64], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.ObjectId = result

	fmt.Printf("%+v\n", object)
}
