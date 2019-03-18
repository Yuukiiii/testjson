package main

import (
	"common/util"
	"fmt"
)


func main() {
	var gcf Config
	if err := util.ParseConf(&gcf, "etc/country_code.yaml"); err != nil {
		return
	}

	var c = gcf.Country
	fmt.Println(c["日本"])
	fmt.Println(c["美国"])

}
