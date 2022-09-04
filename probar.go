package probar

import (
	"fmt"
	"strings"
)

// The first iteration must be with the int 1
func ProgressBar(iteration, total int) {
	var (
		end    = ">"
		length = 35
		fill   = "="
		prefix = "Progress"
		suffix = " Complete"
	)

	percent := float64(iteration) / float64(total) * 100
	filledLength := int(length * iteration / total)

	if iteration == total {
		end = "="
	}

	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
	fmt.Printf("\r%s [%s] %v%%", prefix, bar, int64(percent))
	if iteration == total {
		fmt.Println(suffix)
	}
}
