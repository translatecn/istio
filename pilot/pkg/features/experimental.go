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
	"time"

	"go.uber.org/atomic"

	"istio.io/istio/pkg/env"
	"istio.io/istio/pkg/log"
)

// Define experimental features here.
var (
	SendUnhealthyEndpoints = atomic.NewBool(env.Register(
		"PILOT_SEND_UNHEALTHY_ENDPOINTS",
		false,
		"If enabled, Pilot will include unhealthy endpoints in EDS pushes and even if they are sent Envoy does not use them for load balancing."+
			"  To avoid, sending traffic to non ready endpoints, enabling this flag, disables panic threshold in Envoy i.e. Envoy does not load balance requests"+
			" to unhealthy/non-ready hosts even if the percentage of healthy hosts fall below minimum health percentage(panic threshold).",
	).Get())

	EnablePersistentSessionFilter = atomic.NewBool(env.Register(
		"PILOT_ENABLE_PERSISTENT_SESSION_FILTER",
		false,
		"If enabled, Istiod sets up persistent session filter for listeners, if services have 'PILOT_PERSISTENT_SESSION_LABEL' set.",
	).Get())

	DrainingLabel = env.Register(
		"PILOT_DRAINING_LABEL",
		"istio.io/draining",
		"If not empty, endpoints with the label value present will be sent with status DRAINING.",
	).Get()

	MCSAPIGroup = env.Register("MCS_API_GROUP", "multicluster.x-k8s.io",
		"The group to be used for the Kubernetes Multi-Cluster Services (MCS) API.").Get()

	MCSAPIVersion = env.Register("MCS_API_VERSION", "v1alpha1",
		"The version to be used for the Kubernetes Multi-Cluster Services (MCS) API.").Get()

	EnableMCSServiceDiscovery = env.Register("ENABLE_MCS_SERVICE_DISCOVERY", false, "如果启用，istiod将启用Kubernetes多集群服务（MCS）服务发现模式。在这种模式下，除非通过ServiceExport显式导出，否则集群中的服务端点只能在同一集群中被发现。").Get()
	EnableMCSHost             = env.Register("ENABLE_MCS_HOST", false,
		"如果启用，istiod将为至少一个集群中的每个导出服务（通过ServiceExport）配置Kubernetes多集群服务（MCS）主机（<svc>.<namespace>.svc.clusterset.local）."+
			"但是，客户端必须能够成功地查找这些DNS主机。这意味着必须启用Istio DNS拦截，或者必须使用MCS控制器。要求 ENABLE_MCS_SERVICE_DISCOVERY 也被启用。").Get() &&
		EnableMCSServiceDiscovery

	EnableMCSClusterLocal = env.Register("ENABLE_MCS_CLUSTER_LOCAL", false,
		"如果启用，istiod将处理主机<svc>.<namespace>.svc.cluster。Kubernetes Multi-Cluster Services （MCS）规范中定义的“local”模式。"+
			"“本地”将只路由到与客户端位于同一集群中的那些端点。要求ENABLE_MCS_SERVICE_DISCOVERY和ENABLE_MCS_HOST也被启用。").Get() &&
		EnableMCSHost

	AnalysisInterval = func() time.Duration {
		val, _ := env.Register(
			"PILOT_ANALYSIS_INTERVAL",
			10*time.Second,
			"If analysis is enabled, pilot will run istio analyzers using this value as interval in seconds "+
				"Istio Resources",
		).Lookup()
		if val < 1*time.Second {
			log.Warnf("PILOT_ANALYSIS_INTERVAL %s is too small, it will be set to default 10 seconds", val.String())
			return 10 * time.Second
		}
		return val
	}()

	EnableGatewayAPIDeploymentController = env.Register("PILOT_ENABLE_GATEWAY_API_DEPLOYMENT_CONTROLLER", true,
		"If this is set to true, gateway-api resources will automatically provision in cluster deployment, services, etc").Get()

	EnableGatewayAPIGatewayClassController = env.Register("PILOT_ENABLE_GATEWAY_API_GATEWAYCLASS_CONTROLLER", true,
		"If this is set to true, istiod will create and manage its default GatewayClasses").Get()

	DeltaXds = env.Register("ISTIO_DELTA_XDS", true,
		"If enabled, pilot will only send the delta configs as opposed to the state of the world on a "+
			"Resource Request. This feature uses the delta xds api, but does not currently send the actual deltas.").Get()

	EnableQUICListeners = env.Register("PILOT_ENABLE_QUIC_LISTENERS", false,
		"If true, QUIC listeners will be generated wherever there are listeners terminating TLS on gateways "+
			"if the gateway service exposes a UDP port with the same number (for example 443/TCP and 443/UDP)").Get()

	EnableTLSOnSidecarIngress = env.Register("ENABLE_TLS_ON_SIDECAR_INGRESS", false,
		"If enabled, the TLS configuration on Sidecar.ingress will take effect").Get()

	EnableHCMInternalNetworks = env.Register("ENABLE_HCM_INTERNAL_NETWORKS", false,
		"If enable, endpoints defined in mesh networks will be configured as internal addresses in Http Connection Manager").Get()

	EnableSidecarServiceInboundListenerMerge = env.Register(
		"PILOT_ALLOW_SIDECAR_SERVICE_INBOUND_LISTENER_MERGE",
		false,
		"If set, it allows creating inbound listeners for service ports and sidecar ingress listeners ",
	).Get()

	EnableDualStack = env.RegisterBoolVar("ISTIO_DUAL_STACK", false,
		"If true, Istio will enable the Dual Stack feature.").Get()

	// This is used in injection templates, it is not unused.
	EnableNativeSidecars = env.Register("ENABLE_NATIVE_SIDECARS", false, "If set, used Kubernetes native Sidecar container support. Requires SidecarContainer feature flag.")

	PassthroughTargetPort = env.Register("ENABLE_RESOLUTION_NONE_TARGET_PORT", true,
		"If enabled, targetPort will be supported for resolution=NONE ServiceEntry").Get()

	Enable100ContinueHeaders = env.Register("ENABLE_100_CONTINUE_HEADERS", true,
		"If enabled, istiod will proxy 100-continue headers as is").Get()

	EnableDeferredClusterCreation = env.Register("ENABLE_DEFERRED_CLUSTER_CREATION", true,
		"If enabled, Istio will create clusters only when there are requests. This will save memory and CPU cycles"+
			" in cases where there are lots of inactive clusters and > 1 worker thread").Get()

	EnableDeferredStatsCreation = env.Register("ENABLE_DEFERRED_STATS_CREATION", true,
		"If enabled, Istio will lazily initialize a subset of the stats").Get()

	EnableLocalityWeightedLbConfig = env.Register("ENABLE_LOCALITY_WEIGHTED_LB_CONFIG", false,
		"If enabled, always set LocalityWeightedLbConfig for a cluster, "+
			" otherwise only apply it when locality lb is specified by DestinationRule for a service").Get()

	BypassOverloadManagerForStaticListeners = env.Register("BYPASS_OVERLOAD_MANAGER_FOR_STATIC_LISTENERS", true,
		"If enabled, overload manager will not be applied to static listeners").Get()

	EnableEnhancedDestinationRuleMerge = env.Register("ENABLE_ENHANCED_DESTINATIONRULE_MERGE", true,
		"If enabled, Istio merge destinationrules considering their exportTo fields,"+
			" they will be kept as independent rules if the exportTos are not equal.").Get()

	UnifiedSidecarScoping = env.Register("PILOT_UNIFIED_SIDECAR_SCOPE", true,
		"If true, unified SidecarScope creation will be used. This is only intended as a temporary feature flag for backwards compatibility.").Get()
	EnableEnhancedResourceScoping = env.Register("ENABLE_ENHANCED_RESOURCE_SCOPING", true, "如果启用,meshConfig.discoverySelectors将限制可由pilot处理的CustomResource配置（如Gateway、VirtualService、DestinationRule、Ingress等）。这也将限制根ca证书的分发。").Get()

	EnableGatewayAPI = env.Register("PILOT_ENABLE_GATEWAY_API", true, "If this is set to true, support for Kubernetes gateway-api (github.com/kubernetes-sigs/gateway-api) will be enabled. In addition to this being enabled, the gateway-api CRDs need to be installed.").Get()

	EnableAlphaGatewayAPI = env.Register("PILOT_ENABLE_ALPHA_GATEWAY_API", false, "If this is set to true, support for alpha APIs in the Kubernetes gateway-api (github.com/kubernetes-sigs/gateway-api) will be enabled. In addition to this being enabled, the gateway-api CRDs need to be installed.").Get()

	EnableAnalysis = env.Register("PILOT_ENABLE_ANALYSIS", false, "If enabled, pilot will run istio analyzers and write analysis errors to the Status field of any Istio Resources").Get()

	EnableGatewayAPIStatus = env.Register("PILOT_ENABLE_GATEWAY_API_STATUS", true, "If this is set to true, gateway-api resources will have status written to them").Get()

	EnableLeaderElection = env.Register("ENABLE_LEADER_ELECTION", true, "If enabled (default), starts a leader election client and gains leadership before executing controllers. If false, it assumes that only one instance of istiod is running and skips leader election.").Get()

	EnableMCSAutoExport = env.Register("ENABLE_MCS_AUTO_EXPORT", false, "如果启用，istiod将自动为网格中的每个服务生成Kubernetes Multi-Cluster Services (MCS) ServiceExport资源。在MeshConfig中定义为集群本地的服务被排除在外。").Get()

	FilterGatewayClusterConfig = env.Register("PILOT_FILTER_GATEWAY_CLUSTER_CONFIG", false, "如果启用，Pilot将只发送附加到网关的网关虚拟服务中引用的集群").Get()

	PersistentSessionHeaderLabel = env.Register("PILOT_PERSISTENT_SESSION_HEADER_LABEL", "istio.io/persistent-session-header", "If not empty, services with this label will use header based persistent sessions").Get()
	PersistentSessionLabel       = env.Register("PILOT_PERSISTENT_SESSION_LABEL", "istio.io/persistent-session", "If not empty, services with this label will use cookie based persistent sessions").Get()
)
