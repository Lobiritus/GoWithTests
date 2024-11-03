package main

import (
	"GoWithTests/Clock"
	"os"
	"time"
)

func main() {
	t := time.Now()
	Clock.SVGWriter(os.Stdout, t)
}
