package main

import (
	probar "github.com/seyuta/goprobar"
)

func main() {
	total := 100000

	for i := 1; i <= total; i++ {
		probar.ProgressBar(i, total)
	}
}
