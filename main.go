package main

import (
	"baseservice/testjson/proto/head"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func test(msg proto.Message) {
	fmt.Println((msg).String())
}

func main() {
	h1 := &head.Head{
		MsgTypeName: string("test"),
		Ret:         int32(1),
		Key:         string("key"),
	}
	proto.Marshal(h1)
	fmt.Println(h1)
	test(h1)
}
