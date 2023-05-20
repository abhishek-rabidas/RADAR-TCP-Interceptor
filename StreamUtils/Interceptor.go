package StreamUtils

import "fmt"

type InterceptorDetails struct {
	buffer []byte
}

func InitializeInterceptor(buffer []byte) *InterceptorDetails {
	instance := &InterceptorDetails{
		buffer: buffer,
	}
	instance.GetPayload()
	return instance
}

func (interceptor *InterceptorDetails) GetPayload() {
	fmt.Printf("%x\n", interceptor.buffer)
}
