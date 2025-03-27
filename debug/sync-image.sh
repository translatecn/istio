docker login harbor.ls.com -u admin -p Harbor12345
l() {
	image=$1
	docker pull $image
	kind load docker-image -n koord $image
}

l registry.cn-hangzhou.aliyuncs.com/acejilam/all-in-one:1.58
l registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-details-v1:1.20.2
l registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-productpage-v1:1.20.2
l registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-ratings-v1:1.20.2
l registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-reviews-v1:1.20.2
l registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-reviews-v2:1.20.2
l registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-reviews-v3:1.20.2
l registry.cn-hangzhou.aliyuncs.com/acejilam/grafana:11.2.2-security-01
l registry.cn-hangzhou.aliyuncs.com/acejilam/k8s-sidecar:1.27.5
l registry.cn-hangzhou.aliyuncs.com/acejilam/kiali:v2.0
l registry.cn-hangzhou.aliyuncs.com/acejilam/loki:3.2.0
l registry.cn-hangzhou.aliyuncs.com/acejilam/mygo:v1.24.0
l registry.cn-hangzhou.aliyuncs.com/acejilam/prometheus-config-reloader:v0.76.0
l registry.cn-hangzhou.aliyuncs.com/acejilam/prometheus:v2.54.1
l registry.cn-hangzhou.aliyuncs.com/acejilam/proxyv2:1.24.3
l registry.cn-hangzhou.aliyuncs.com/acejilam/skywalking-oap-server:9.7.0
l registry.cn-hangzhou.aliyuncs.com/acejilam/skywalking-ui:9.1.0
