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
		stepSizeInBytes := int(stepSize)
		memorySink = append(memorySink, make([]byte, stepSizeInBytes))
		end := time.Now().UnixNano()
		durationForMemoryIncrease := end - start
		sleepDuration := time.Duration(int64(stepDuration) - durationForMemoryIncrease)
		_, _ = fmt.Fprintf(os.Stderr, "\r Swallowed up %s from virtual address space            ", bytesize.ByteSize(len(memorySink)*stepSizeInBytes))
		time.Sleep(sleepDuration)
	}
}
