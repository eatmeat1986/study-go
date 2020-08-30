package main

import (
	"fmt"
	"time"
)

func main() {
	timeChan := time.After(10 * time.Second)

	<-timeChan
	fmt.Println("finished")
}
