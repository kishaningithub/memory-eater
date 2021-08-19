package memoryeater

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"os"
	"time"
)

func Eat(stepSize bytesize.ByteSize, stepDuration time.Duration) {
	var memorySink [][]byte
	for {
		start := time.Now().UnixNano()
		memorySink = append(memorySink, make([]byte, int(stepSize)))
		end := time.Now().UnixNano()
		durationForMemoryIncrease := end - start
		sleepDuration := time.Duration(int64(stepDuration) - durationForMemoryIncrease)
		_, _ = fmt.Fprintf(os.Stderr, "\r Swallowed up %s              ", bytesize.ByteSize(len(memorySink)*int(stepSize)))
		time.Sleep(sleepDuration)
	}
}
