package testpkg

import (
	"fmt"
	"regexp"
)

//TestRegexp xxx
func TestRegexp() {
	testStr := "peach punch"
	regBin, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Printf("regBin:%v\n", regBin)

	fmt.Printf("1:%v\n", regBin.Match([]byte(testStr)))
	fmt.Printf("2:%v\n", regBin.MatchString(testStr))

	fmt.Printf("3:%v\n", regBin.FindStringIndex(testStr))
	fmt.Printf("4:%v\n", regBin.FindString(testStr))
	fmt.Printf("5:%v\n", regBin.FindAllString(testStr, -1))

	fmt.Printf("6:%v\n", regBin.FindStringSubmatchIndex(testStr))
	fmt.Printf("7:%v\n", regBin.FindStringSubmatch(testStr))
	fmt.Printf("8:%v\n", regBin.FindAllStringSubmatch(testStr, -1))
	fmt.Printf("9:%v\n", regBin.FindAllStringSubmatchIndex(testStr, -1))

	fmt.Printf("10:%v\n", regBin.ReplaceAllString(testStr, "fruit"))
}
