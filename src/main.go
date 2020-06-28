package main

import (
	"fmt"
	"mypkg"

	"github.com/Unknwon/goconfig"
)

func main() {
	confFile, _ := goconfig.LoadConfigFile("test.conf")
	tmpConfVal, _ := confFile.GetValue("aaa", "conf_name1")
	fmt.Println(tmpConfVal)
	mypkg.TestFunc()
}
