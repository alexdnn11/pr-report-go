package main

import (
	"fmt"
	"./wps"
)

func main() {
wpRes := wps.getWP()
fmt.Print(wpRes)
}
