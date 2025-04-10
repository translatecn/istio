#!/usr/bin/env zsh
export KUBECONFIG=/Users/acejilam/.kube/koord
kubectl get virtualservices.networking.istio.io -n test -oyaml | yq '.items[0].spec' | drop-age.py
set -x
kubectl get deployment -n test | drop-age.py
kubectl get hpa -n test | drop-age.py
kubectl get svc -n test --show-labels | drop-age.py
kubectl get endpoints -n test | drop-age.py
kubectl get pods -n test --show-labels | drop-age.py
kubectl get canaries.flagger.app -n test | drop-age.py

#k -n test scale deployment/podinfo --replicas=2
#k -n test set image deployment/podinfo podinfod=ccr.ccs.tencentyun.com/acejilam/podinfo:6.0.1
