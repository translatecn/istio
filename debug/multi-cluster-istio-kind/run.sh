set -ex
cd kind-create
bash ./create-cluster.sh
bash ./create-cluster.sh
bash ./install-metallb.sh
bash ./install-cacerts.sh
cd -
cd istio-create
bash ./install-istio.sh
bash ./enable-endpoint-discovery.sh
cd -
cd testing
bash ./deploy-app.sh
cd -

kubectl --context cluster1 -n sample exec -it deployment/helloworld-v2 -- bash -c 'while true; do curl http://helloworld.sample:5000/hello; done'
