package util

import (
	"bufio"
	"log"
	"os"
	"os/exec"
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

// TODO: ensure this can be run on any OS
func IsAppleComputer() bool {
	cmd := exec.Command("uname", "-a")
	out, err := cmd.Output()
	if err != nil {
		return false
	}

	if string(out) == "Darwin" {
		return true
	}
	return false
}
