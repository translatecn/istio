package discovery

import (
	"os"

	"istio.io/istio/pilot/pkg/model"
)

func init() {
	os.Setenv("REVISION", "default")
	os.Setenv("PILOT_CERT_PROVIDER", "istiod")
	os.Setenv("POD_NAME", "istiod-6fd85f74ff-xtqr6")
	os.Setenv("POD_NAMESPACE", "istio-system")
	os.Setenv("SERVICE_ACCOUNT", "istiod")
	os.Setenv("KUBECONFIG", "/var/run/secrets/remote/config")
	os.Setenv("CA_TRUSTED_NODE_ACCOUNTS", "istio-system/ztunnel")
	os.Setenv("PILOT_TRACE_SAMPLING", "100")
	os.Setenv("PILOT_ENABLE_ANALYSIS", "false")
	os.Setenv("CLUSTER_ID", "Kubernetes")
	os.Setenv("GOMEMLIMIT", "0")
	os.Setenv("GOMAXPROCS", "1")
	model.JwksExtraRootCABundlePath = "/cacerts/extra.pem" //

	os.Setenv("PILOT_ENABLE_ANALYSIS", "true")
	os.Setenv("PILOT_ENABLE_IP_AUTOALLOCATE", "true")
	os.Setenv("PILOT_ENABLE_NODE_UNTAINT_CONTROLLERS", "true")

	os.Setenv("ENABLE_MCS_AUTO_EXPORT", "true")
	os.Setenv("ENABLE_MCS_SERVICE_DISCOVERY", "true")
	os.Setenv("ENABLE_MCS_HOST", "true")
}

// configmap  istio-ca-root-cert  	-> /var/run/secrets/istiod/ca
// secret     cacerts    			-> /etc/cacerts
// secret     istio-kubeconfig    	-> /var/run/secrets/remote
// secret     istiod-tls    		-> /var/run/secrets/istiod/tls

// istio-csr-ca-configmap
