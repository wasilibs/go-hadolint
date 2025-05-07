# go-hadolint

go-hadolint is a distribution of [hadolint][1], that can be built with Go. It does not actually reimplement any
functionality of hadolint in Go, instead compiling it with the GHC WASI backend, and
executing with the pure Go Wasm runtime [wazero][2]. This means that `go install` or `go run`
can be used to execute it, with no need to rely on separate package managers such as pnpm,
on any platform that Go supports.

## Behavior Differences

- Output formatters for Codeclimate are not supported

## Installation

Precompiled binaries are available in the [releases](https://github.com/wasilibs/go-hadolint/releases).
Alternatively, install the plugin you want using `go install`.

```bash
$ go install github.com/wasilibs/go-hadolint/cmd/hadolint@latest
```

To avoid installation entirely, it can be convenient to use `go run`

```bash
$ go run github.com/wasilibs/go-hadolint/cmd/hadolint@latest *.sh
```

_Due to [potential build breakage](https://github.com/golang/go/issues/71192) unrelated to this project,
`go tool` is not supported._

Note that due to the sandboxing of the filesystem when using Wasm, currently only files that descend
from the current directory when executing the tool are accessible to it, i.e., `../docker/Dockerfile` or
`/separate/root/Dockerfile` will not be found.

[1]: https://github.com/hadolint/hadolint
[2]: https://wazero.io/
