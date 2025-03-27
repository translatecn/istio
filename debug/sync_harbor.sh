docker login harbor.ls.com -u admin -p Harbor12345

t() {
	image=$1
	new_name=$(echo $1 | sed 's#registry.cn-hangzhou.aliyuncs.com#harbor.ls.com#g')
	echo $new_name
	until docker pull $image; do
		echo "docker pull $image, retrying in 5 seconds..."
		sleep 5
	done

	docker tag $image $new_name

	until docker push $new_name; do
		echo "docker push $new_name, retrying in 5 seconds..."
		sleep 5
	done

}

t registry.cn-hangzhou.aliyuncs.com/acejilam/k8s-sidecar:1.27.5
t registry.cn-hangzhou.aliyuncs.com/acejilam/kiali:v2.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/loki:3.2.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/proxyv2:1.24.3
t registry.cn-hangzhou.aliyuncs.com/acejilam/pilot:1.24.3
t registry.cn-hangzhou.aliyuncs.com/acejilam/install-cni:1.24.3-distroless
t registry.cn-hangzhou.aliyuncs.com/acejilam/pilot:1.24.3-distroless
t registry.cn-hangzhou.aliyuncs.com/acejilam/proxyv2:1.24.3-distroless
t registry.cn-hangzhou.aliyuncs.com/acejilam/ztunnel:1.24.3-distroless
t registry.cn-hangzhou.aliyuncs.com/acejilam/frr:9.1.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/metallb_controller:v0.14.9
t registry.cn-hangzhou.aliyuncs.com/acejilam/metallb_speaker:v0.14.9
t registry.cn-hangzhou.aliyuncs.com/acejilam/flagger:1.40.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/flagger-loadtester:0.35.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/metrics-server:v0.6.3
t registry.cn-hangzhou.aliyuncs.com/acejilam/prometheus-config-reloader:v0.76.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/prometheus:v2.54.1

t registry.cn-hangzhou.aliyuncs.com/acejilam/skywalking-oap-server:9.7.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/skywalking-ui:9.1.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/podinfo:6.0.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/podinfo:6.0.1
t registry.cn-hangzhou.aliyuncs.com/acejilam/podinfo:6.0.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/all-in-one:1.58
t registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-details-v1:1.20.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-productpage-v1:1.20.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-ratings-v1:1.20.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-reviews-v1:1.20.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-reviews-v2:1.20.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/examples-bookinfo-reviews-v3:1.20.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/grafana:11.2.2-security-01
t registry.cn-hangzhou.aliyuncs.com/acejilam/mygo:v1.24.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/nginx:1.14.2
t registry.cn-hangzhou.aliyuncs.com/acejilam/dnsutils:1.3
t registry.cn-hangzhou.aliyuncs.com/acejilam/centos:7
t registry.cn-hangzhou.aliyuncs.com/acejilam/go-httpbin:v2.15.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/curl:latest
t registry.cn-hangzhou.aliyuncs.com/acejilam/buildpack-deps:24.04
