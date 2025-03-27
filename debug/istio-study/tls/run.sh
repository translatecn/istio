#!/usr/bin/env bash
kubectl delete ns foo
kubectl delete ns bar
kubectl delete ns legacy

kubectl create ns foo
kubectl apply -f <(istioctl kube-inject -f httpbin.yaml) -n foo
kubectl apply -f <(istioctl kube-inject -f sleep.yaml) -n foo
kubectl create ns bar
kubectl apply -f <(istioctl kube-inject -f httpbin.yaml) -n bar
kubectl apply -f <(istioctl kube-inject -f sleep.yaml) -n bar

kubectl create ns legacy
kubectl apply -f httpbin.yaml -n legacy
kubectl apply -f sleep.yaml -n legacy

#kubectl apply -f peer-strict-global.yaml
kubectl apply -f peer-strict-ns.yaml

kubectl wait --for=condition=Ready pods --timeout=3000s --all

kubectl exec "$(kubectl get pod -l app=sleep -n bar -o jsonpath={.items..metadata.name})" -c sleep -n bar -- curl http://httpbin.foo:8000/ip -s -o /dev/null -w "%{http_code}\n"

for from in "foo" "bar" "legacy"; do
	for to in "foo" "bar" "legacy"; do
		echo ${from} "--->" ${to}
		podName=$(kubectl get pod -l app=sleep -n ${from} -o jsonpath={.items..metadata.name})
		kubectl exec "${podName}" -c sleep -n ${from} -- curl -s "http://httpbin.${to}:8000/ip" -s -o /dev/null -w "sleep.${from} to httpbin.${to}: %{http_code}\n"
	done
done

kubectl exec "$(kubectl get pod -l app=sleep -n foo -o jsonpath={.items..metadata.name})" -c sleep -n foo -- curl -s http://httpbin.foo:8000/headers -s
