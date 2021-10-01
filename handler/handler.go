package handler

import (
	"fmt"
	"time"
)

func AsyncHandler() {
	time.Sleep(8 * time.Second)
	fmt.Println("I wait 8 seconds for a response")
}
