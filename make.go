// +build ignore
package main

import (
	"log"
	"os/exec"
)

func main() {
	targs := []string{
		"github.com/rene00/khaos/server/khaosd",
	}
	baseArgs := []string{"install", "-v"}
	args := append(baseArgs, targs...)
	cmd := exec.Command("go", args...)
	log.Printf("Running go %q", args)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error building main binaries: %v\n", err)
	}
}
