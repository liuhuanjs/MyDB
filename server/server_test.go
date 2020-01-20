package server

import (
	"fmt"
	"testing"
	"time"
)

func TestForLoop(t *testing.T) {
	i := 1
	for {
		fmt.Printf("count : %d\n", i)
		i++
		time.Sleep(1 * time.Second)
	}
}
