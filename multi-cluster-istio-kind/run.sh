set -ex

rm /Users/acejilam/.kube/koord || echo skip

cd kind-setup
bash ./create-cluster.sh
bash ./install-metallb.sh
bash ./install-cacerts.sh
cd -
cd istio-setup
bash ./install-istio.sh
bash ./enable-endpoint-discovery.sh
cd -
cd testing
bash ./deploy-application.sh
cd -

kubectl --context cluster1 -n sample exec -it deployment/helloworld-v2 -- bash -c 'while true; do curl http://helloworld.sample:5000/hello; done'
