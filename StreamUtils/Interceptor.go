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

	if !interceptor.sync {
		_, flag := strings.CutPrefix(payload, "02ff")

		if flag {
			return
		}

		_, flag = strings.CutPrefix(payload, "0734")

		if flag {
			return
		}
	}

	fmt.Println(payload)
	fmt.Println()

}
