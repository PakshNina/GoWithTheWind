package main

import (
	"fmt"
	"strings"
	"time"
)

func regularStringBuilder() string {
	var res string
	for i := 0; i < 500000; i++ {
		res += "x"
	}
	return res
}

func fastStringBuilder() string {
	var sb strings.Builder
	for i := 0; i < 500000; i++ {
		sb.WriteString("x")
	}
	return sb.String()
}

func main() {
	timeStart := time.Now()
	regularStringBuilder()
	fmt.Println(time.Now().Sub(timeStart).Seconds())

	timeStart = time.Now()
	fastStringBuilder()
	fmt.Println(time.Now().Sub(timeStart).Seconds())
}
