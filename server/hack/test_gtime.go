package main
import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)
func main() {
	fmt.Printf("%T %v\n", gtime.M, gtime.M)
	fmt.Printf("%T %v\n", time.Minute, time.Minute)
}