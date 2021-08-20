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
		bytes := make([]byte, stepSizeInBytes)
		for i := 0; i < stepSizeInBytes; i++ {
			bytes[i] = '2'
		}
		memorySink = append(memorySink, bytes)
		end := time.Now().UnixNano()
		durationForMemoryIncrease := end - start
		sleepDuration := time.Duration(int64(stepDuration) - durationForMemoryIncrease)
		_, _ = fmt.Fprintf(os.Stderr, "\r Swallowed up %s from RAM (RSS Space)            ", bytesize.ByteSize(len(memorySink)*stepSizeInBytes))
		time.Sleep(sleepDuration)
	}
}
