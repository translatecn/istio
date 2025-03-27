// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package features

import (
	"strings"
	"time"

	"istio.io/istio/pkg/config/constants"
	"istio.io/istio/pkg/env"
	"istio.io/istio/pkg/jwt"
	"istio.io/istio/pkg/util/sets"
)

var (
	// HTTP10 will add "accept_http_10" to http outbound listeners. Can also be set only for specific sidecars via meta.
	HTTP10 = env.Register(
		"PILOT_HTTP10",
		false,
		"Enables the use of HTTP 1.0 in the outbound HTTP listeners, to support legacy applications.",
	).Get()

	ScopeGatewayToNamespace = env.Register(
		"PILOT_SCOPE_GATEWAY_TO_NAMESPACE",
		false,
		"If enabled, a gateway workload can only select gateway resources in the same namespace. "+
			"Gateways with same selectors in different namespaces will not be applicable.",
	).Get()

	JwksFetchMode = func() jwt.JwksFetchMode {
		v := env.Register(
			"PILOT_JWT_ENABLE_REMOTE_JWKS",
			"false",
			"Mode of fetching JWKs from JwksUri in RequestAuthentication. Supported value: "+
				"istiod, false, hybrid, true, envoy. The client fetching JWKs is as following: "+
				"istiod/false - Istiod; hybrid/true - Envoy and fallback to Istiod if JWKs server is external; "+
				"envoy - Envoy.",
		).Get()
		return jwt.ConvertToJwksFetchMode(v)
	}()

	// IstiodServiceCustomHost allow user to bring a custom address or multiple custom addresses for istiod server
	// for examples: 1. istiod.mycompany.com  2. istiod.mycompany.com,istiod-canary.mycompany.com
	IstiodServiceCustomHost = env.Register("ISTIOD_CUSTOM_HOST", "",
		"Custom host name of istiod that istiod signs the server cert. "+
			"Multiple custom host names are supported, and multiple values are separated by commas.").Get()

	EnableServiceEntrySelectPods = env.Register("PILOT_ENABLE_SERVICEENTRY_SELECT_PODS", true,
		"If enabled, service entries with selectors will select pods from the cluster. "+
			"It is safe to disable it if you are quite sure you don't need this feature").Get()

	ValidationWebhookConfigName = env.Register("VALIDATION_WEBHOOK_CONFIG_NAME", "istio-istio-system",
		"If not empty, the controller will automatically patch validatingwebhookconfiguration when the CA certificate changes. "+
			"Only works in kubernetes environment.").Get()

	DisableMxALPN = env.Register("PILOT_DISABLE_MX_ALPN", false,
		"If true, pilot will not put istio-peer-exchange ALPN into TLS handshake configuration.",
	).Get()

	ALPNFilter = env.Register("PILOT_ENABLE_ALPN_FILTER", true,
		"If true, pilot will add Istio ALPN filters, required for proper protocol sniffing.",
	).Get()

	WorkloadEntryAutoRegistration = env.Register("PILOT_ENABLE_WORKLOAD_ENTRY_AUTOREGISTRATION", true,
		"Enables auto-registering WorkloadEntries based on associated WorkloadGroups upon XDS connection by the workload.").Get()

	WorkloadEntryCleanupGracePeriod = env.Register("PILOT_WORKLOAD_ENTRY_GRACE_PERIOD", 10*time.Second,
		"The amount of time an auto-registered workload can remain disconnected from all Pilot instances before the "+
			"associated WorkloadEntry is cleaned up.").Get()

	WorkloadEntryHealthChecks = env.Register("PILOT_ENABLE_WORKLOAD_ENTRY_HEALTHCHECKS", true,
		"Enables automatic health checks of WorkloadEntries based on the config provided in the associated WorkloadGroup").Get()

	WorkloadEntryCrossCluster = env.Register("PILOT_ENABLE_CROSS_CLUSTER_WORKLOAD_ENTRY", true, "If enabled, pilot will read WorkloadEntry from other clusters, selectable by Services in that cluster.").Get()

	WasmRemoteLoadConversion = env.Register("ISTIO_AGENT_ENABLE_WASM_REMOTE_LOAD_CONVERSION", true,
		"If enabled, Istio agent will intercept ECDS resource update, downloads Wasm module, "+
			"and replaces Wasm module remote load with downloaded local module file.").Get()

	PilotJwtPubKeyRefreshInterval = env.Register(
		"PILOT_JWT_PUB_KEY_REFRESH_INTERVAL",
		20*time.Minute,
		"The interval for istiod to fetch the jwks_uri for the jwks public key.",
	).Get()

	// EnableUnsafeAssertions enables runtime checks to test assertions in our code. This should never be enabled in
	// production; when assertions fail Istio will panic.

	// EnableUnsafeDeltaTest enables runtime checks to test Delta XDS efficiency. This should never be enabled in
	// production.

	EnableEnvoyFilterMetrics = env.Register("PILOT_ENVOY_FILTER_STATS", false,
		"If true, Pilot will collect metrics for envoy filter operations.").Get()

	EnableRouteCollapse = env.Register("PILOT_ENABLE_ROUTE_COLLAPSE_OPTIMIZATION", true,
		"If true, Pilot will merge virtual hosts with the same routes into a single virtual host, as an optimization.").Get()

	MulticlusterHeadlessEnabled = env.Register("ENABLE_MULTICLUSTER_HEADLESS", true,
		"If true, the DNS name table for a headless service will resolve to same-network endpoints in any cluster.").Get()

	InsecureKubeConfigOptions = func() sets.String {
		v := env.Register(
			"PILOT_INSECURE_MULTICLUSTER_KUBECONFIG_OPTIONS",
			"",
			"Comma separated list of potentially insecure kubeconfig authentication options that are allowed for multicluster authentication."+
				"Support values: all authProviders (`gcp`, `azure`, `exec`, `openstack`), "+
				"`clientKey`, `clientCertificate`, `tokenFile`, and `exec`.").Get()
		return sets.New(strings.Split(v, ",")...)
	}()

	CanonicalServiceForMeshExternalServiceEntry = env.Register("LABEL_CANONICAL_SERVICES_FOR_MESH_EXTERNAL_SERVICE_ENTRIES", false,
		"If enabled, metadata representing canonical services for ServiceEntry resources with a location of mesh_external will be populated"+
			"in the cluster metadata for those endpoints.").Get()

	// This is a feature flag, can be removed if protobuf proves universally better.
	KubernetesClientContentType = env.Register("ISTIO_KUBE_CLIENT_CONTENT_TYPE", "protobuf",
		"The content type to use for Kubernetes clients. Defaults to protobuf. Valid options: [protobuf, json]").Get()

	JwksResolverInsecureSkipVerify = env.Register("JWKS_RESOLVER_INSECURE_SKIP_VERIFY", false,
		"If enabled, istiod will skip verifying the certificate of the JWKS server.").Get()

	EnableSelectorBasedK8sGatewayPolicy = env.Register("ENABLE_SELECTOR_BASED_K8S_GATEWAY_POLICY", true,
		"If disabled, Gateway API gateways will ignore workloadSelector policies, only"+
			"applying policies that select the gateway with a targetRef.").Get()

	// Useful for IPv6-only EKS clusters. See https://aws.github.io/aws-eks-best-practices/networking/ipv6/ why it assigns an additional IPv4 NAT address.
	// Also see https://github.com/istio/istio/issues/46719 why this flag is required
	EnableAdditionalIpv4OutboundListenerForIpv6Only = env.RegisterBoolVar("ISTIO_ENABLE_IPV4_OUTBOUND_LISTENER_FOR_IPV6_CLUSTERS", false,
		"If true, pilot will configure an additional IPv4 listener for outbound traffic in IPv6 only clusters, e.g. AWS EKS IPv6 only clusters.").Get()

	EnableAutoSni = env.Register("ENABLE_AUTO_SNI", true,
		"If enabled, automatically set SNI when `DestinationRules` do not specify the same").Get()

	EnableVtprotobuf = env.Register("ENABLE_VTPROTOBUF", true,
		"If true, will use optimized vtprotobuf based marshaling. Requires a build with -tags=vtprotobuf.").Get()

	GatewayAPIDefaultGatewayClass = env.Register("PILOT_GATEWAY_API_DEFAULT_GATEWAYCLASS_NAME", "istio",
		"Name of the default GatewayClass").Get()

	ManagedGatewayController = env.Register("PILOT_GATEWAY_API_CONTROLLER_NAME", "istio.io/gateway-controller",
		"Gateway API controller name. istiod will only reconcile Gateway API resources referencing a GatewayClass with this controller name").Get()

	EnableInboundRetryPolicy = env.Register("ENABLE_INBOUND_RETRY_POLICY", true,
		"If true, enables retry policy for inbound routes which automatically retries requests that were reset before it reaches the service.").Get()

	Exclude503FromDefaultRetries = env.Register("EXCLUDE_UNSAFE_503_FROM_DEFAULT_RETRY", true,
		"If true, excludes unsafe retry on 503 from default retry policy.").Get()

	PreferDestinationRulesTLSForExternalServices = env.Register("PREFER_DESTINATIONRULE_TLS_FOR_EXTERNAL_SERVICES", true,
		"If true, external services will prefer the TLS settings from DestinationRules over the metadata TLS settings.").Get()
	EnableDebugOnHTTP         = env.Register("ENABLE_DEBUG_ON_HTTP", true, "If this is set to false, the debug interface will not be enabled, recommended for production").Get()
	SharedMeshConfig          = env.Register("SHARED_MESH_CONFIG", "", "为共享的MeshConfig设置加载额外的配置映射。标准网格配置将优先考虑。").Get()
	EnableCAServer            = env.Register("ENABLE_CA_SERVER", true, "If this is set to false, will not create CA server in istiod.").Get()
	PilotCertProvider         = env.Register("PILOT_CERT_PROVIDER", constants.CertProviderIstiod, "The provider of Pilot DNS certificate. K8S RA will be used for k8s.io/NAME. 'istiod' value will sign using Istio build in CA. Other values will not not generate TLS certs, but still distribute .ca. Only used if custom certificates are not mounted.").Get()
	LocalClusterSecretWatcher = env.Register("LOCAL_CLUSTER_SECRET_WATCHER", false, "如果启用，集群秘密监视程序将监视外部集群的命名空间，而不是配置集群").Get()
	ExternalIstiod            = env.Register("EXTERNAL_ISTIOD", false, "如果设置为true，一个Istiod将控制包括CA在内的远程集群。").Get()

	EnableNodeUntaintControllers = env.Register("PILOT_ENABLE_NODE_UNTAINT_CONTROLLERS", false, "如果启用，将运行带有cni pod的控制器。如果您禁用了环境初始化容器，则应该启用此选项。").Get()
	EnableIPAutoallocate         = env.Register("PILOT_ENABLE_IP_AUTOALLOCATE", false, "如果启用，pilot将启动一个控制器，该控制器为没有用户提供IP的ServiceEntry分配IP地址。当与DNS捕获结合使用时，它允许发送到ServiceEntry的流量的tcp路由。").Get()
	MultiRootMesh                = env.Register("ISTIO_MULTIROOT_MESH", false, "如果启用，mesh将支持由多个ISTIO_MUTUAL mTLS的trustchor签名的证书").Get()

	InjectionWebhookConfigName = env.Register("INJECTION_WEBHOOK_CONFIG_NAME", "istio-sidecar-injector", "Name of the mutatingwebhookconfiguration to patch, if istioctl is not used.").Get()
	EnableUnsafeAdminEndpoints = env.Register("UNSAFE_ENABLE_ADMIN_ENDPOINTS", false, "如果将此设置为true，则将在调试界面上暴露危险的管理端点。不建议用于生产。").Get()
	EnableUnsafeAssertions     = env.Register("UNSAFE_PILOT_ENABLE_RUNTIME_ASSERTIONS", false, "如果启用，将执行附加运行时断言。这些检查既昂贵又失败。因此，这应该只用于测试。").Get()
	EnableUnsafeDeltaTest      = env.Register("UNSAFE_PILOT_ENABLE_DELTA_TEST", false, "如果启用，则会添加Delta XDS效率的附加运行时测试。这些检查非常昂贵，因此应该只用于测试，而不是用于生产。").Get()

	ResolveHostnameGateways = env.Register("RESOLVE_HOSTNAME_GATEWAYS", true, "如果为true，则服务的LoadBalancer地址中的主机名将在控制平面解析，以便在跨网络网关中使用。").Get()
	ClusterName             = env.Register("CLUSTER_ID", constants.DefaultClusterName, "Defines the cluster and service registry that this Istiod instance belongs to").Get()

	EnableK8SServiceSelectWorkloadEntries = env.RegisterBoolVar("PILOT_ENABLE_K8S_SELECT_WORKLOAD_ENTRIES", true, "如果启用，带选择器的Kubernetes服务将选择具有匹配标签的工作负载条目。如果您非常确定不需要此功能，则禁用它是安全的").Get()

	RemoteClusterTimeout   = env.Register("PILOT_REMOTE_CLUSTER_TIMEOUT", 30*time.Second, "After this timeout expires, pilot can become ready without syncing data from clusters added via remote-secrets. Setting the timeout to 0 disables this behavior.").Get()
	MultiNetworkGatewayAPI = env.Register("PILOT_MULTI_NETWORK_DISCOVER_GATEWAY_API", true, "If true, Pilot will discover labeled Kubernetes gateway objects as multi-network gateways.").Get()

	InformerWatchNamespace        = env.Register("ISTIO_WATCH_NAMESPACE", "", "If set, limit Kubernetes watches to a single namespace. Warning: only a single namespace can be set.").Get()
	ValidateWorkloadEntryIdentity = env.Register("ISTIO_WORKLOAD_ENTRY_VALIDATE_IDENTITY", true, "如果启用，将验证工作负载的标识是否与它所关联的WorkloadEntry的标识匹配，以便进行运行状况检查和自动注册。此标志仅为向后兼容性而添加，并将在将来的版本中删除").Get()
)

func UnsafeFeaturesEnabled() bool {
	return EnableUnsafeAdminEndpoints || EnableUnsafeAssertions || EnableUnsafeDeltaTest
}
