package codec

import "io"

type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string
}

type Codec interface {
	io.Closer                         // releasing resources
	ReadHeader(*Header) error         // read header
	ReadBody(interface{}) error       // read body
	Write(*Header, interface{}) error // write data
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	JsonType Type = "application/json"
	GobType  Type = "application/gob"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[JsonType] = NewJsonCodec
	NewCodecFuncMap[GobType] = NewGobCodec
}
