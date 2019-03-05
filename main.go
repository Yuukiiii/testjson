package main

import (
	"baseservice/testjson/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	h1 := &head.Head{
		MsgTypeName: string("test"),
		Ret:         int32(1),
		Key:         string("key"),
	}
	proto.Marshal(h1)
	fmt.Println(h1.String())

}
