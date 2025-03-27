FROM registry.cn-hangzhou.aliyuncs.com/acejilam/tinygo AS build
WORKDIR /src

COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN tinygo build -o main.wasm -scheduler=none -target=wasi ./main.go

FROM scratch

COPY --from=build /src/main.wasm ./plugin.wasm
