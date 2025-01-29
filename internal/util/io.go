package util

import (
	"bufio"
	"log"
	"os"
)

func ReadStdin() *[]string {
	stat, _ := os.Stdin.Stat()
	var list []string
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			list = append(list, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Failed to read from stdin: %v", err)
		}
	}
	return &list
}
