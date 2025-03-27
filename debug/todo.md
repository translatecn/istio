
- XDSServer ConfigUpdate 更新时机
    - MutatingWebhookConfiguration 发生变化                         Full: true,    Reason: model.NewReasonStats(model.TagUpdate)
    - istio-gateway-status-leader  成为leader                      Full: true,    Reason: model.NewReasonStats(model.GlobalUpdate),
    - secret 有变化  全部(!helm,!service-account-token)          Full:            false,
                                                               ConfigsUpdated:  sets.New(model.ConfigKey{Kind: kind.Secret, Name: name, Namespace: namespace}),
                                                               Reason:          model.NewReasonStats(model.SecretTrigger),
    - istio configmap 发生变化                                   Full: true,     Reason: model.NewReasonStats(model.GlobalUpdate),
    - 证书发生变化                                                Full: true,     Reason: model.NewReasonStats(model.GlobalUpdate),
    - Service 发生变化                  			                Full:           true,
                                    			                ConfigsUpdated: sets.New(model.ConfigKey{Kind: kind.ServiceEntry, Name: string(curr.Hostname), Namespace: curr.Attributes.Namespace}),
                                    			                Reason:         model.NewReasonStats(model.ServiceUpdate),
    - PilotGatewayAPI 大部分CRD 发生更新 				            Full:           true,
                                    				            ConfigsUpdated: sets.New(model.ConfigKey{Kind: kind.MustFromGVK(curr.GroupVersionKind), Name: curr.Name, Namespace: curr.Namespace}),
                                    				            Reason:         model.NewReasonStats(model.ConfigUpdate),
    - namespace 发生变化，且与istio mesh 满足                      Full:   true,
                                                         		Reason: model.NewReasonStats(model.NamespaceUpdate),
    - secret 全部(!helm,!service-account-token),满足gw映射        Full: true,
                                                             	ConfigsUpdated: map[model.ConfigKey]struct{}{
                                                             		{
                                                             			Kind:      kind.KubernetesGateway,
                                                             			Name:      gw.Name,
                                                             			Namespace: gw.Namespace,
                                                             		}: {},
                                                             	},
                                                             	Reason: model.NewReasonStats(model.SecretTrigger),
    - gateway 发生更新                                           Full: true, Reason: NewReasonStats(NetworksTrigger)
    - endpoint 发生更新                                          Full:           pushType == model.FullPush,
                                                                ConfigsUpdated: sets.New(model.ConfigKey{Kind: kind.ServiceEntry, Name: serviceName, Namespace: namespace}),
                                                                Reason:         model.NewReasonStats(model.EndpointUpdate),
    - jwt 公钥发生变化                                            Full: true, Reason: model.NewReasonStats(model.UnknownTrigger)


- 证书更新的时机   什么时候调用 UpdateTrustAnchor
- Service更新的时机   什么时候调用 NotifyServiceHandlers
- gateway 更新的时机
    (istio mesh.meshNetworks 发生更新; ./etc/istio/config/meshNetworks) -> SetNetworks  <... reloadGateways
    /etc/resolv.conf         NetworkGatewaysHandler.NotifyGatewayHandlers() <... reloadGateways            // 1、监听的域名定时更新
    2?         NetworkGatewaysHandler.NotifyGatewayHandlers() <... reloadGateways


.handler()
handlers {
AddRunFunction
WithReconciler(?
type ComponentBuilder interface {
type handler interface {
type registerDependency interface {
type Informer[T controllers.Object] interface {
    AddEventHandler
type MulticlusterController interface {
type Injector interface {
type Watcher interface {
type NetworksWatcher interface {
type XDSUpdater interface {

proxy.istio.io/overrides  注入时，如有已有则覆盖



createNewContext




reloadGateways: (NotifyGatewayHandlers)
- refreshAndNotify 域名解析定时更新




sidecar、router、waypoint、ztunnel


DestinationRule、
ServiceEntry
getSidecarScope
getPodByProxy



func (s *DiscoveryServer) Start(stopCh <-chan struct{}) {
	go s.WorkloadEntryController.Run(stopCh)
	go s.handleUpdates(stopCh) // ✅
	go s.periodicRefreshMetrics(stopCh)
	go s.sendPushes(stopCh) // ✅
	go s.Cache.Run(stopCh)
}

