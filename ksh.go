package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	//"log"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("ksh> ")
		if !scanner.Scan() {
			break
		}

		trimmed := strings.TrimSpace(scanner.Text())
		if trimmed == "" {
			continue
		}

		split := strings.Fields(trimmed)
		if split[0] == "exit" {
				return
		}
		cmd := exec.Command(split[0], split[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
		fmt.Printf("command '%s' not found\n but can be installed using sudo install '%s'\n", split[0], split[0])
		}
	}
}
