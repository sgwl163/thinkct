package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"net/http"
	"time"
)

func main() {
	client := g.Client().SetTimeout(5 * gtime.M)
	// We can't easily inspect the internal transport of g.Client(), but we can see its behavior.
	fmt.Printf("Timeout: %v\n", 5*gtime.M)
}