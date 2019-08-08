package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Microsecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	//ch := make(chan string)
	ch := make(chan string, numOfRunner) //使用bufferchannel
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(FirstResponse())
	//t.Log(time.Second * 1)
	time.Sleep(time.Second * 1)
	t.Log("After: ", runtime.NumGoroutine())
}
