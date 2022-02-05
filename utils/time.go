package utils
import (
	"time"
)
func Delay(durationInSec int) {
	time.Sleep(time.Duration(durationInSec)* time.Second)
}