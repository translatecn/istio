package agent

import "os"

func init() {
	os.Setenv("CA_ADDR", "istiod.istio-system.svc:15012")
	os.Setenv("GOMAXPROCS", "")
	os.Setenv("GOMEMLIMIT", "")
	os.Setenv("HOST_IP", "192.168.33.13")
	os.Setenv("ISTIO_CPU_LIMIT", "")
	os.Setenv("ISTIO_META_APP_CONTAINERS", "reviews")
	os.Setenv("ISTIO_META_CLUSTER_ID", "Kubernetes")
	os.Setenv("ISTIO_META_INTERCEPTION_MODE", "REDIRECT")
	os.Setenv("ISTIO_META_MESH_ID", "cluster.local")
	os.Setenv("ISTIO_META_NODE_NAME", "vm2404")
	os.Setenv("ISTIO_META_OWNER", "kubernetes://apis/apps/v1/namespaces/default/deployments/reviews-v1")
	os.Setenv("ISTIO_META_POD_PORTS", `[ {"containerPort":9080,"protocol":"TCP"} ]`)
	os.Setenv("ISTIO_META_WORKLOAD_NAME", "reviews-v1")
	os.Setenv("PILOT_CERT_PROVIDER", "istiod")
	os.Setenv("POD_NAME", "")
	os.Setenv("POD_NAMESPACE", "istio-system")
	os.Setenv("PROXY_CONFIG", "{}")
	os.Setenv("SERVICE_ACCOUNT", "default")
	os.Setenv("TRUST_DOMAIN", "cluster.local")
}
