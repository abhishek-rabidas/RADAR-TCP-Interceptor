package entity

type Stream struct {
	Id      []byte
	Length  []byte
	Payload []byte
}

type ObjectData struct {
	X        float64
	Y        float64
	XSpeed   float64
	YSpeed   float64
	Length   float64
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
