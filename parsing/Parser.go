package parsing

import (
	"encoding/hex"
	"fmt"
	"radar/parsing/entity"
	"strconv"
	"strings"
)

func ParseStreams(streamBytes string) []entity.Stream {

	//i := 0

	var output []entity.Stream

	/*	for {

		if i+4 >= len(streamBytes) {
			break
		}

		var stream entity.Stream = entity.Stream{}

		stream.Id = streamBytes[i : i+4]
		i += 4

		stream.Length, _ = strconv.Atoi(streamBytes[i : i+2])
		i += 2

		j := i + (stream.Length * 2)
		stream.Payload = convertToBinary(convertFormat(streamBytes[i:j])) //reversing the payload to make it big endian and then getting the binary representation
		i = j

		//fmt.Printf("%+v\n", stream)

		output = append(output, stream)

		if stream.Id == "0502" {
			ParseObjectData(stream.Payload)
		} else if stream.Id == "0501" {
			ParseObjectControlData(stream.Payload)
		}

	}*/

	return output
}

func ParseObjectData(payload string) {
	var object entity.ObjectData = entity.ObjectData{}

	//x_coordinate
	result, err := strconv.ParseUint(payload[1:14], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.X = result

	//y_coordinate
	result, err = strconv.ParseUint(payload[14:27], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Y = result

	//X_speed
	result, err = strconv.ParseUint(payload[27:38], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.XSpeed = result

	//Y_speed
	result, err = strconv.ParseUint(payload[38:49], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.YSpeed = result

	//Length
	result, err = strconv.ParseUint(payload[49:56], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.Length = result

	//ObjectId
	result, err = strconv.ParseUint(payload[56:64], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.ObjectId = result

	fmt.Printf("%+v\n", object)

}

func ParseObjectControlData(payload string) {
	var object entity.ObjectControl = entity.ObjectControl{}

	result, err := strconv.ParseUint(payload[0:7], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.NumberOfObjects = result

	result, err = strconv.ParseUint(payload[8:15], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.NumberOfMessages = result

	result, err = strconv.ParseUint(payload[16:23], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.CycleDuration = result

	//Length
	result, err = strconv.ParseUint(payload[28:31], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.ObjectDataFormat = result

	result, err = strconv.ParseUint(payload[32:63], 2, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	object.CycleCount = result

	fmt.Printf("%+v\n", object)

}

func convertFormat(payload string) string {
	var out []byte

	for i := len(payload) - 1; i >= 0; i-- {
		out = append(out, payload[i])
	}

	for i := 0; i < len(out); i += 2 {
		temp := out[i]
		out[i] = out[i+1]
		out[i+1] = temp
	}

	return string(out)
}

func convertToBinary(payload string) string {
	decodedBytes, err := hex.DecodeString(payload)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return ""
	}

	// Convert decoded bytes to string

	binary := fmt.Sprintf("%08b", decodedBytes)

	binary = strings.ReplaceAll(binary, "[", "")
	binary = strings.ReplaceAll(binary, "]", "")
	binary = strings.ReplaceAll(binary, " ", "")

	//fmt.Println(binary)
	return binary
}
