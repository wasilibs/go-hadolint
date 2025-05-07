package main

import (
	"os"

	"github.com/wasilibs/go-hadolint/internal/runner"
	"github.com/wasilibs/go-hadolint/internal/wasm"
)

func main() {
	os.Exit(runner.Run("hadolint", os.Args[1:], wasm.Hadolint, os.Stdin, os.Stdout, os.Stderr, "."))
}
