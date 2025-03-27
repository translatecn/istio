set -ex
#docker-install-harbor.sh
#sleep 3
#curl -k -u "admin:Harbor12345" -X POST -H "Content-Type: application/json" "https://harbor.ls.com/api/v2.0/projects/" -d '{"project_name": "acejilam", "public": true}'
#./debug/sync_harbor.sh &
install-k8s-by-kind.sh koord v1.28.15
k8s-use-ls-harbor.py
docker-install-metallb.sh
docker-install-metrics-server.sh
install-istio.py harbor
docker-install-flagger.sh
