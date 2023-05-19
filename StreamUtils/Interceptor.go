package StreamUtils

type InterceptorDetails struct {
	startChecksum []byte
	endChecksum   []byte
	buffer        []byte
}

func InitializeInterceptor(startChecksum []byte, endChecksum []byte, buffer []byte) *InterceptorDetails {
	return &InterceptorDetails{
		startChecksum: startChecksum,
		endChecksum:   endChecksum,
		buffer:        buffer,
	}
}

func (interceptor *InterceptorDetails) GetPayload() {

}
