package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "hadolint",
		LibraryRepo: "hadolint/hadolint",
		GoReleaser:  true,
	})
	boot.Main()
}
