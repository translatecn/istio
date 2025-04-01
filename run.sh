set -ex
docker-install-harbor.sh
sleep 3
./debug/sync_harbor.sh

export HOST_IP=$(print-proxy.py | grep -v unset | grep -v export)

install-k8s-by-kind.sh koord v1.28.15
k8s-use-ls-harbor.py
docker-install-metallb.sh
docker-install-metrics-server.sh
install-istio.py harbor example
docker-install-flagger.sh
