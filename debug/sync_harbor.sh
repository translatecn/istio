#echo "${HOST_IP} harbor.ls.com" >/etc/hosts
#skopeo login -u admin harbor.ls.com -p Harbor12345 --tls-verify=false


set -ex

t() {
	old_image=$1
	new_name=$(echo $1 | sed 's#ccr.ccs.tencentyun.com#harbor.ls.com#g')
#	skopeo copy --override-os linux --insecure-policy docker://${old_image} docker://${new_name} --src-tls-verify=false --dest-tls-verify=false

	echo $new_name
	until docker pull $old_image; do
		echo "docker pull $old_image, retrying in 5 seconds..."
		sleep 5
	done

	docker tag $old_image $new_name

	until docker push $new_name; do
		echo "docker push $new_name, retrying in 5 seconds..."
		sleep 5
	done
}

t ccr.ccs.tencentyun.com/acejilam/ubuntu:22.04
t ccr.ccs.tencentyun.com/acejilam/frr:9.1.0
t ccr.ccs.tencentyun.com/acejilam/metallb_controller:v0.14.9
t ccr.ccs.tencentyun.com/acejilam/metallb_speaker:v0.14.9
t ccr.ccs.tencentyun.com/acejilam/nginx:1.14.2

t ccr.ccs.tencentyun.com/acejilam/k8s-sidecar:1.27.5
t ccr.ccs.tencentyun.com/acejilam/kiali:v2.0
t ccr.ccs.tencentyun.com/acejilam/loki:3.2.0
t ccr.ccs.tencentyun.com/acejilam/install-cni:1.24.3-distroless
t ccr.ccs.tencentyun.com/acejilam/pilot:1.24.3
t ccr.ccs.tencentyun.com/acejilam/pilot:1.24.3-distroless
t ccr.ccs.tencentyun.com/acejilam/proxyv2:1.24.3
t ccr.ccs.tencentyun.com/acejilam/proxyv2:1.24.3-distroless
t ccr.ccs.tencentyun.com/acejilam/ztunnel:1.24.3-distroless
t ccr.ccs.tencentyun.com/acejilam/flagger:1.40.0
t ccr.ccs.tencentyun.com/acejilam/flagger-loadtester:0.35.0
t ccr.ccs.tencentyun.com/acejilam/metrics-server:v0.6.3
t ccr.ccs.tencentyun.com/acejilam/metrics-server:v0.7.2
t ccr.ccs.tencentyun.com/acejilam/prometheus-config-reloader:v0.76.0
t ccr.ccs.tencentyun.com/acejilam/prometheus:v2.54.1
t ccr.ccs.tencentyun.com/acejilam/skywalking-oap-server:9.7.0
t ccr.ccs.tencentyun.com/acejilam/skywalking-ui:9.1.0
t ccr.ccs.tencentyun.com/acejilam/podinfo:6.0.0
t ccr.ccs.tencentyun.com/acejilam/podinfo:6.0.1
t ccr.ccs.tencentyun.com/acejilam/podinfo:6.0.2
t ccr.ccs.tencentyun.com/acejilam/all-in-one:1.58
t ccr.ccs.tencentyun.com/acejilam/examples-bookinfo-details-v1:1.20.2
t ccr.ccs.tencentyun.com/acejilam/examples-bookinfo-productpage-v1:1.20.2
t ccr.ccs.tencentyun.com/acejilam/examples-bookinfo-ratings-v1:1.20.2
t ccr.ccs.tencentyun.com/acejilam/examples-bookinfo-reviews-v1:1.20.2
t ccr.ccs.tencentyun.com/acejilam/examples-bookinfo-reviews-v2:1.20.2
t ccr.ccs.tencentyun.com/acejilam/examples-bookinfo-reviews-v3:1.20.2
t ccr.ccs.tencentyun.com/acejilam/grafana:11.2.2-security-01
t ccr.ccs.tencentyun.com/acejilam/mygo:v1.24.0
t ccr.ccs.tencentyun.com/acejilam/dnsutils:1.3
t ccr.ccs.tencentyun.com/acejilam/centos:7
t ccr.ccs.tencentyun.com/acejilam/go-httpbin:v2.15.0
t ccr.ccs.tencentyun.com/acejilam/curl:latest
t ccr.ccs.tencentyun.com/acejilam/buildpack-deps:24.04

t ccr.ccs.tencentyun.com/acejilam/microservices-demo-adservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-cartservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-checkoutservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-currencyservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-emailservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-frontend:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/busybox:latest
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-loadgenerator:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-paymentservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-productcatalogservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-recommendationservice:v0.9.0
t ccr.ccs.tencentyun.com/acejilam/redis:alpine
t ccr.ccs.tencentyun.com/acejilam/microservices-demo-shippingservice:v0.9.0
