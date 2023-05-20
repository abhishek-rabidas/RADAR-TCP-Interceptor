package StreamUtils

import (
	"encoding/hex"
	"fmt"
	"radar/config"
	"strings"
)

type InterceptorDetails struct {
	buffer []byte
	config *config.Interceptor
}

func InitializeInterceptor(buffer []byte, interceptor *config.Interceptor) *InterceptorDetails {
	instance := &InterceptorDetails{
		buffer: buffer,
		config: interceptor,
	}

	return instance
}

func (interceptor *InterceptorDetails) GetPayload() {
	var isFound bool

	startChecksum := strings.Join(interceptor.config.StartChecksumHex, "")
	endChecksum := strings.Join(interceptor.config.EndChecksumHex, "")

	var payload string = hex.EncodeToString(interceptor.buffer)

	//pruning the payload to get the actual payload by removing start checksum, end checksum and trailing bits (0)

	payload = strings.TrimRight(payload, "0")

	payload, isFound = strings.CutPrefix(payload, startChecksum)
	if !isFound {
		return
	}

	payload, isFound = strings.CutSuffix(payload, endChecksum)
	if !isFound {
		return
	}

	fmt.Println(payload)
	fmt.Println()

}
