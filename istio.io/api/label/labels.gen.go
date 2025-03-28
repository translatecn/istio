// GENERATED FILE -- DO NOT EDIT

package label

type FeatureStatus int

const (
	Alpha FeatureStatus = iota
	Beta
	Stable
)

func (s FeatureStatus) String() string {
	switch s {
	case Alpha:
		return "Alpha"
	case Beta:
		return "Beta"
	case Stable:
		return "Stable"
	}
	return "Unknown"
}

type ResourceTypes int

const (
	Unknown ResourceTypes = iota
	Any
	Deployment
	Gateway
	GatewayClass
	Namespace
	Node
	Pod
	Service
	ServiceAccount
	ServiceEntry
	WorkloadEntry
)

func (r ResourceTypes) String() string {
	switch r {
	case 1:
		return "Any"
	case 2:
		return "Deployment"
	case 3:
		return "Gateway"
	case 4:
		return "GatewayClass"
	case 5:
		return "Namespace"
	case 6:
		return "Node"
	case 7:
		return "Pod"
	case 8:
		return "Service"
	case 9:
		return "ServiceAccount"
	case 10:
		return "ServiceEntry"
	case 11:
		return "WorkloadEntry"
	}
	return "Unknown"
}

// Instance describes a single resource label
type Instance struct {
	// The name of the label.
	Name string

	// Description of the label.
	Description string

	// FeatureStatus of this label.
	FeatureStatus FeatureStatus

	// Hide the existence of this label when outputting usage information.
	Hidden bool

	// Mark this label as deprecated when generating usage information.
	Deprecated bool

	// The types of resources this label applies to.
	Resources []ResourceTypes
}

var (
	GatewayManaged = Instance{
		Name: "gateway.istio.io/managed",
		Description: "Automatically added to all resources [automatically " +
			"created](/docs/tasks/traffic-management/ingress/gateway-api/#automated-deployment) " +
			"by Istio Gateway controller, to indicate which controller " +
			"created the resource. Users should not set this label " +
			"themselves.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			ServiceAccount,
			Deployment,
			Service,
		},
	}

	IoK8sNetworkingGatewayGatewayName = Instance{
		Name: "gateway.networking.k8s.io/gateway-name",
		Description: "Automatically added to all resources [automatically " +
			"created](/docs/tasks/traffic-management/ingress/gateway-api/#automated-deployment) " +
			"by Istio Gateway controller to indicate which `Gateway` " +
			"resulted in the object creation. Users should not set " +
			"this label themselves.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			ServiceAccount,
			Deployment,
			Service,
		},
	}

	IoIstioDataplaneMode = Instance{
		Name: "istio.io/dataplane-mode",
		Description: `When set on a resource, indicates the [data plane mode](/docs/overview/dataplane-modes/) to use.
Possible values: "ambient", "none".
Note: users wishing to use sidecar mode should see the "istio-injection" label; there is no value on this label to configure sidecars.
`,
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
			Namespace,
		},
	}

	IoIstioRev = Instance{
		Name: "istio.io/rev",
		Description: "Istio control plane revision associated with the " +
			"resource; e.g. `canary`",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
		},
	}

	IoIstioTag = Instance{
		Name: "istio.io/tag",
		Description: "Istio control plane tag name associated with the " +
			"resource; e.g. `canary`",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
		},
	}

	IoIstioUseWaypoint = Instance{
		Name: "istio.io/use-waypoint",
		Description: `When set on a resource, indicates the resource has an associated waypoint with the given name.
The waypoint is assumed to be in the same namespace; for cross-namespace, see "istio.io/use-waypoint-namespace".

When set or a "Pod" or a "Service", this binds that specific resource to the waypoint.
When set on a "Namespace", this applies to all "Pod"/"Service" in the namespace.

Note: the waypoint must allow the type, see "stio.io/waypoint-for".
`,
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
			WorkloadEntry,
			Service,
			ServiceEntry,
			Namespace,
		},
	}

	IoIstioUseWaypointNamespace = Instance{
		Name: "istio.io/use-waypoint-namespace",
		Description: `When set on a resource, indicates the resource has an associated waypoint in the provided namespace.
This must be set in addition to "istio.io/use-waypoint", when a cross-namespace reference is desired.
`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
			WorkloadEntry,
			Service,
			ServiceEntry,
			Namespace,
		},
	}

	IoIstioWaypointFor = Instance{
		Name: "istio.io/waypoint-for",
		Description: `When set on a waypoint (either by its specific "Gateway", or for the entire collection on the "GatewayClass"),
indicates the type of traffic this waypoint can handle.

Valid options: "service", "workload", "all", and "none".
`,
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			GatewayClass,
			Gateway,
		},
	}

	NetworkingEnableAutoallocateIp = Instance{
		Name: "networking.istio.io/enable-autoallocate-ip",
		Description: `Configures whether a "ServiceEntry" without any "spec.addresses" set should get an IP address automatically allocated for it.

Valid options: "true", "false"
`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			ServiceEntry,
		},
	}

	NetworkingGatewayPort = Instance{
		Name: "networking.istio.io/gatewayPort",
		Description: "IstioGatewayPortLabel overrides the default 15443 value " +
			"to use for a multi-network gateway's port",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
		},
	}

	OperatorComponent = Instance{
		Name: "operator.istio.io/component",
		Description: "Istio operator component name of the resource, e.g. " +
			"`Pilot`",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorManaged = Instance{
		Name: "operator.istio.io/managed",
		Description: "Set to `Reconcile` if the Istio operator will reconcile " +
			"the resource.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorVersion = Instance{
		Name: "operator.istio.io/version",
		Description: "The Istio operator version that installed the resource, " +
			"e.g. `1.6.0`",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	SecurityTlsMode = Instance{
		Name: "security.istio.io/tlsMode",
		Description: "Specifies the TLS mode supported by a sidecar proxy. " +
			"Valid values are 'istio', 'disabled'. When injecting " +
			"sidecars into Pods, the sidecar injector will set the " +
			"value of this label to 'istio' indicating that the " +
			"sidecar is capable of supporting mTLS. Clients injected " +
			"with sidecar proxies will opportunistically use this " +
			"label to determine whether or not to secure the traffic " +
			"to this workload using Istio mutual TLS.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ServiceCanonicalName = Instance{
		Name:          "service.istio.io/canonical-name",
		Description:   "The name of the canonical service a workload belongs to",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ServiceCanonicalRevision = Instance{
		Name: "service.istio.io/canonical-revision",
		Description: "The name of a revision within a canonical service that " +
			"the workload belongs to",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarInject = Instance{
		Name: "sidecar.istio.io/inject",
		Description: "Specifies whether or not an Envoy sidecar should be " +
			"automatically injected into the workload.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	TopologyCluster = Instance{
		Name: "topology.istio.io/cluster",
		Description: "This label is applied to a workload internally that " +
			"identifies the Kubernetes cluster containing the " +
			"workload. The cluster ID is specified during Istio " +
			"installation for each cluster via " +
			"`values.global.multiCluster.clusterName`. It should be " +
			"noted that this is only used internally within Istio and " +
			"is not an actual label on workload pods. If a pod " +
			"contains this label, it will be overridden by Istio " +
			"internally with the cluster ID specified during Istio " +
			"installation. This label provides a way to select " +
			"workloads by cluster when using DestinationRules. For " +
			"example, a service owner could create a DestinationRule " +
			"containing a subset per cluster and then use these " +
			"subsets to control traffic flow to each cluster " +
			"independently.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	TopologyNetwork = Instance{
		Name: "topology.istio.io/network",
		Description: `用于标识一个或多个 Pod 所在网络的标签。此标签由 Istio 在内部使用，用于将处于同一 L3 域/网络中的 Pod 进行分组。
Istio 假定处于同一网络中的 Pod 相互直接可达。当 Pod 处于不同网络时，通常会使用 Istio 网关（例如东向/西向网关）来建立连接（采用 AUTO_PASSTHROUGH 模式）。此标签可应用于以下资源，以帮助实现 Istio 多网络配置的自动化。

* Istio 系统命名空间：将此标签应用于系统命名空间可为由控制平面管理的 Pod 建立默认网络。这通常在控制平面安装期间使用管理员指定的值进行配置。

* Pod：将此标签应用于 Pod 可以在每个 Pod 的层面上覆盖默认网络设置。这通常通过 Webhook 注入的方式应用于 Pod，但服务所有者也可以在 Pod 上手动指定该设置。每个集群中的 Istio 安装使用管理员指定的值来配置 Webhook 注入。

* 网关服务：将此标签应用于 Istio 网关所服务的项目，表明 Istio 在配置跨网络流量时应将此服务用作网络的网关。Istio 会将位于网络之外的 Pod 配置为通过“spec.externalIPs”、“status.loadBalancer.ingress[].ip”（对于 NodePort 服务，为 Node 的地址）来访问网关服务。当安装网关（例如东向网关）时配置此标签，并且应与控制平面的默认网络（如 Istio 系统命名空间标签指定的那样）或目标 Pod 的网络相匹配。`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
			Pod,
			Service,
		},
	}

	TopologySubzone = Instance{
		Name: "topology.istio.io/subzone",
		Description: "User-provided node label for identifying the locality " +
			"subzone of a workload. This allows admins to specify a " +
			"more granular level of locality than what is offered by " +
			"default with Kubernetes regions and zones.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Node,
		},
	}
)

func AllResourceLabels() []*Instance {
	return []*Instance{
		&GatewayManaged,
		&IoK8sNetworkingGatewayGatewayName,
		&IoIstioDataplaneMode,
		&IoIstioRev,
		&IoIstioTag,
		&IoIstioUseWaypoint,
		&IoIstioUseWaypointNamespace,
		&IoIstioWaypointFor,
		&NetworkingEnableAutoallocateIp,
		&NetworkingGatewayPort,
		&OperatorComponent,
		&OperatorManaged,
		&OperatorVersion,
		&SecurityTlsMode,
		&ServiceCanonicalName,
		&ServiceCanonicalRevision,
		&SidecarInject,
		&TopologyCluster,
		&TopologyNetwork,
		&TopologySubzone,
	}
}

func AllResourceTypes() []string {
	return []string{
		"Any",
		"Deployment",
		"Gateway",
		"GatewayClass",
		"Namespace",
		"Node",
		"Pod",
		"Service",
		"ServiceAccount",
		"ServiceEntry",
		"WorkloadEntry",
	}
}
