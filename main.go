package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"protobuf/head"

)


func main()  {
	tcpd := Tcpd_Rest{
		Url:	"test",
		Id:		"1",
	}

	str, err := json.Marshal(tcpd)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(string(str))
	}

	tcpd2 := Tcpd_Rest{}
	err2 := json.Unmarshal([]byte(string(str)), &tcpd2)

	if (err2 != nil) {
		fmt.Println(err2)
	} else {
		fmt.Println(tcpd2.Url)
	}
}
