docker rm whoami1 whoami2 --force
docker run -d -p 8001:80 --name whoami1 ccr.ccs.tencentyun.com/acejilam/whoami:v1.10.1 --verbose
docker run -d -p 8002:80 --name whoami2 ccr.ccs.tencentyun.com/acejilam/whoami:v1.10.1 --verbose

export W1=$(ipconfig getifaddr en0)
yq -i '.spec.address = env(W1) ' ./WorkloadEntry1.yaml
yq -i '.spec.address = env(W1) ' ./WorkloadEntry2.yaml

kubectl apply -f ./DestinationRule.yaml

# curl whoami.internal.com 可以通过
kubectl apply -f ./WorkloadEntry2.yaml -f ./WorkloadEntry1.yaml -f ServiceEntry.yaml
kubectl -n default exec -it deployment/istio-test -- bash -c 'curl whoami.internal.com '

# curl whoami.com 可以通过
kubectl apply -f ./VirtualService.yaml
kubectl -n default exec -it deployment/istio-test -- bash -c 'curl whoami.com '

# curl ingress 可以通过
Iip=$(kubectl -n istio-system get svc istio-ingressgateway -o json | jq '.status.loadBalancer.ingress[0].ip' | tr -d '\"')
kubectl apply -f ./Gateway.yaml
kubectl -n default exec -it deployment/istio-test -- bash -c "curl -v -H 'Host: whoami.com' http://$Iip"

kubectl -n default exec -it deployment/istio-test -- bash -c 'for i in $(seq 1 1000) ;do curl whoami.com ;done'
