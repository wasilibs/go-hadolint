FROM debian:12-slim

RUN apt-get update && apt-get install -y binaryen build-essential curl git jq unzip xz-utils zstd

WORKDIR /ghc
RUN curl https://gitlab.haskell.org/haskell-wasm/ghc-wasm-meta/-/raw/master/bootstrap.sh | sh

WORKDIR /workspace
ADD buildtools/wasm/version.txt version.txt
ADD buildtools/wasm/patch.txt patch.txt
RUN curl -L https://github.com/hadolint/hadolint/archive/28fc39f92e2768d8cc9f94d6e2fc16bd4397f092.tar.gz | tar -xz --strip-components 1 -C /workspace
RUN patch -p1 < patch.txt

RUN . ~/.ghc-wasm/env && wasm32-wasi-cabal update && wasm32-wasi-cabal build --allow-newer --dependencies-only
RUN . ~/.ghc-wasm/env && wasm32-wasi-cabal build --ghc-options='-single-threaded -with-rtsopts="-A4m"' --allow-newer exe:hadolint

RUN wasm-opt -o hadolint.wasm --flatten --rereloop --converge -O3 dist-newstyle/build/wasm32-wasi/ghc-*/hadolint-*/x/hadolint/build/hadolint/hadolint.wasm

CMD ["cp", "hadolint.wasm", "/out/"]
