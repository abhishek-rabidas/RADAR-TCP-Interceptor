package StreamUtils

import (
	"bytes"
	"radar/config"
	"radar/parsing"
	"strconv"
)

type InterceptorDetails struct {
	buffer []byte
	config *config.Interceptor
	sync   bool
}

func InitializeInterceptor(buffer []byte, interceptor *config.Interceptor, sync bool) *InterceptorDetails {
	instance := &InterceptorDetails{
		buffer: buffer,
		config: interceptor,
		sync:   sync,
	}

	return instance
}

func (interceptor *InterceptorDetails) GetPayload() {
	var isFound bool

	var startChecksumByte = make([]byte, 4)
	var endChecksumByte = make([]byte, 4)

	for idx, checksum := range interceptor.config.StartChecksumHex {
		decodeString, _ := strconv.ParseUint(checksum, 16, 64)

		startChecksumByte[idx] = byte(decodeString)
	}

	for idx, checksum := range interceptor.config.EndChecksumHex {
		decodeString, _ := strconv.ParseUint(checksum, 16, 64)

		endChecksumByte[idx] = byte(decodeString)
	}

	interceptor.buffer = bytes.TrimRight(interceptor.buffer, "\x00")

	interceptor.buffer, isFound = bytes.CutPrefix(interceptor.buffer, startChecksumByte)
	if !isFound {
		return
	}

	interceptor.buffer, isFound = bytes.CutSuffix(interceptor.buffer, endChecksumByte)
	if !isFound {
		return
	}

	parsing.ParseByteStreams(interceptor.buffer)

	//fmt.Printf("%x\n", interceptor.buffer)

}
