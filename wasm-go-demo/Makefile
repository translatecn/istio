
build:
	tinygo build -o main.wasm -scheduler=none -target=wasi main.go
	docker build -t registry.cn-hangzhou.aliyuncs.com/acejilam/test:wasm .
	docker push registry.cn-hangzhou.aliyuncs.com/acejilam/test:wasm
#	kubectl create configmap new-filter --from-file=new-filter.wasm=main.wasm
#	kubectl apply -f ef
