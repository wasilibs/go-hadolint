FROM debian:12-slim

RUN apt-get update && apt-get install -y binaryen build-essential curl git jq unzip xz-utils zstd

WORKDIR /ghc
# Use version before tail call extension usage
RUN curl -f -L --retry 5 https://gitlab.haskell.org/haskell-wasm/ghc-wasm-meta/-/archive/ada3b8fa0f763e4dccb2b1f6bbf2518bff2a7c6e/ghc-wasm-meta-ada3b8fa0f763e4dccb2b1f6bbf2518bff2a7c6e.tar.gz | tar xz --strip-components=1
RUN ./setup.sh

WORKDIR /workspace
ADD buildtools/wasm/version.txt version.txt
ADD buildtools/wasm/patch.txt patch.txt
RUN curl -L https://github.com/hadolint/hadolint/archive/28fc39f92e2768d8cc9f94d6e2fc16bd4397f092.tar.gz | tar -xz --strip-components 1 -C /workspace
RUN patch -p1 < patch.txt

RUN . ~/.ghc-wasm/env && wasm32-wasi-cabal update && wasm32-wasi-cabal build --allow-newer exe:hadolint

RUN wasm-opt -o hadolint.wasm --flatten --rereloop --converge -O3 dist-newstyle/build/wasm32-wasi/ghc-*/hadolint-*/x/hadolint/build/hadolint/hadolint.wasm

CMD ["cp", "hadolint.wasm", "/out/"]
