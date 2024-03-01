package axrpc

import (
	"encoding/json"
	"fmt"
	"github.com/axliupore/axrpc/codec"
	"log"
	"net"
	"testing"
	"time"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	Accept(l)
}

func TestServer(t *testing.T) {
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple axrpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	_ = json.NewEncoder(conn).Encode(DefaultOption)
	c := codec.NewJsonCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = c.Write(h, fmt.Sprintf("axrpc req %d", h.Seq))
		_ = c.ReadHeader(h)
		var reply string
		_ = c.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
