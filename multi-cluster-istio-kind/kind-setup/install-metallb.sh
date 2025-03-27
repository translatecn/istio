#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

NUM_CLUSTERS="${NUM_CLUSTERS:-2}"

for i in $(seq "${NUM_CLUSTERS}"); do
	echo "Starting metallb deployment in cluster${i}"
	kubectl apply -f ./metallb-native.yaml --context "cluster${i}"
	sleep 5
	kubectl wait --for=condition=Ready pod --all -A --timeout=3000s --context "cluster${i}"
	kubectl apply -f ./metallb-cr-${i}.yaml --context "cluster${i}"
	echo "----"
done
