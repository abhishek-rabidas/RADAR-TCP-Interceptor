package entity

type Stream struct {
	Id      []byte
	Length  []byte
	Payload []byte
}

type ObjectData struct {
	X        uint64
	Y        uint64
	XSpeed   uint64
	YSpeed   uint64
	Length   uint64
	ObjectId uint64
}

type ObjectData2 struct {
	X        string
	Y        string
	XSpeed   string
	YSpeed   string
	Length   string
	ObjectId string
}

type ObjectControl struct {
	NumberOfObjects  uint64
	NumberOfMessages uint64
	CycleDuration    uint64
	ObjectDataFormat uint64
	CycleCount       uint64
}
