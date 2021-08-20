package memoryeater

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"os"
	"time"
)

func Eat(stepSize bytesize.ByteSize, stepDuration time.Duration) {
	leak := make(map[int][]byte)
	ticker := time.NewTicker(stepDuration)
	i := 0
	for range ticker.C {
		stepSizeInBytes := int(stepSize)
		bytes := make([]byte, stepSizeInBytes)
		for j := 0; j < stepSizeInBytes; j++ {
			bytes[j] = '2' // Using to move data from VSZ to RSS in order to eat up the RAM
		}
		leak[i] = bytes
		i++
		_, _ = fmt.Fprintf(os.Stderr, "\r Swallowed up %s from RAM (RSS Space)            ", bytesize.ByteSize(len(leak)*stepSizeInBytes))
	}
}
