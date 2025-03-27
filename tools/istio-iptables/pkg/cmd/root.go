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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"istio.io/istio/pkg/flag"
	"istio.io/istio/pkg/log"
	"istio.io/istio/tools/istio-iptables/pkg/capture_over"
	"istio.io/istio/tools/istio-iptables/pkg/config"
	"istio.io/istio/tools/istio-iptables/pkg/constants"
	dep "istio.io/istio/tools/istio-iptables/pkg/dependencies"
	"istio.io/istio/tools/istio-iptables/pkg/validation_over"
)

const InvalidDropByIptables = "INVALID_DROP"

func handleErrorWithCode(err error, code int) {
	log.Error(err)
	os.Exit(code)
}

func bindCmdlineFlags(cfg *config.Config, cmd *cobra.Command) {
	fs := cmd.Flags()
	fmt.Println(os.Environ())
	flag.Bind(fs, constants.EnvoyPort, "p", "指定将所有TCP流量重定向到的特使端口。", &cfg.ProxyPort)
	flag.BindEnv(fs, constants.InboundCapturePort, "z", "所有到pod/VM的入站TCP流量都应该重定向到该端口。", &cfg.InboundCapturePort)
	flag.BindEnv(fs, constants.InboundTunnelPort, "e", "指定入方向tcp流量的istio tunnel端口。", &cfg.InboundTunnelPort)
	flag.BindEnv(fs, constants.ProxyUID, "u", "指定不应用重定向的用户的UID。通常，这是代理容器的UID。", &cfg.ProxyUID)
	flag.BindEnv(fs, constants.ProxyGID, "g", "指定不应用重定向的用户的GID（与-u param的默认值相同）。", &cfg.ProxyGID)
	flag.BindEnv(fs, constants.InboundInterceptionMode, "m", "用于将入站连接重定向到Envoy的模式，可以是“redirect”或“TPROXY”。", &cfg.InboundInterceptionMode)
	flag.BindEnv(fs, constants.InboundPorts, "b", "以逗号分隔的入站端口列表，流量将被重定向到Envoy（可选）。通配符\\“*\\”可用于配置所有端口的重定向。空列表将禁用。", &cfg.InboundPortsInclude)
	flag.BindEnv(fs, constants.LocalExcludePorts, "d", "从重定向到Envoy时要排除的入站端口列表（可选），以逗号分隔。仅适用于所有入站流量（即\\“*\\”）被重定向时。", &cfg.InboundPortsExclude)
	flag.BindEnv(fs, constants.ExcludeInterfaces, "c", "逗号分隔的网卡列表（可选）。既不会捕获入站流量，也不会捕获出站流量。", &cfg.ExcludeInterfaces)
	flag.BindEnv(fs, constants.ServiceCidr, "i", "逗号分隔的CIDR形式的IP范围列表，以重定向到envoy（可选）。通配符“*\\”可用于重定向所有出站流量。空列表将禁用所有出站。", &cfg.OutboundIPRangesInclude)
	flag.BindEnv(fs, constants.ServiceExcludeCidr, "x", "以逗号分隔的CIDR形式不允许重定向的IP范围列表。仅适用于所有出站流量(即“*”)正在被重定向。", &cfg.OutboundIPRangesExclude)
	flag.BindEnv(fs, constants.OutboundPorts, "q", "为重定向到Envoy而显式包含的出站端口列表，以逗号分隔。", &cfg.OutboundPortsInclude)
	flag.BindEnv(fs, constants.LocalOutboundPortsExclude, "o", "要排除重定向到Envoy的出站端口列表，以逗号分隔。", &cfg.OutboundPortsExclude)
	flag.BindEnv(fs, constants.KubeVirtInterfaces, "k", "以逗号分隔的虚拟接口列表，这些虚拟接口的入方向流量（来自VM）将被视为出方向流量。", &cfg.KubeVirtInterfaces)
	flag.BindEnv(fs, constants.InboundTProxyMark, "t", "", &cfg.InboundTProxyMark)
	flag.BindEnv(fs, constants.InboundTProxyRouteTable, "r", "", &cfg.InboundTProxyRouteTable)
	flag.BindEnv(fs, constants.DryRun, "n", "不要调用任何外部依赖，比如iptables。", &cfg.DryRun)
	flag.BindEnv(fs, constants.TraceLogging, "", "使用LOG链为每个iptables规则插入跟踪日志。", &cfg.TraceLogging)
	flag.BindEnv(fs, constants.IptablesProbePort, "", "设置故障检测监听端口。", &cfg.IptablesProbePort)
	flag.BindEnv(fs, constants.ProbeTimeout, "", "故障检测超时。", &cfg.ProbeTimeout)
	flag.BindEnv(fs, constants.SkipRuleApply, "", "跳过iptables apply。", &cfg.SkipRuleApply)
	flag.BindEnv(fs, constants.RunValidation, "", "iptables进行验证。", &cfg.RunValidation)
	flag.BindEnv(fs, constants.RedirectDNS, "", "启用istio-agent捕获dns流量功能。", &cfg.RedirectDNS)
	// Allow binding to a different var, for consistency with other components
	flag.AdditionalEnv(fs, constants.RedirectDNS, "ISTIO_META_DNS_CAPTURE")
	flag.BindEnv(fs, constants.DropInvalid, "", "在iptables规则中启用无效删除。", &cfg.DropInvalid)
	// This could have just used the default but for backwards compat we support the old env.
	flag.AdditionalEnv(fs, constants.DropInvalid, InvalidDropByIptables)
	flag.BindEnv(fs, constants.DualStack, "", "启用ipv4/ipv6双栈重定向。", &cfg.DualStack)
	// Allow binding to a different var, for consistency with other components
	flag.AdditionalEnv(fs, constants.DualStack, "ISTIO_DUAL_STACK")
	flag.BindEnv(fs, constants.CaptureAllDNS, "", "不是只捕获到DNS服务器IP的DNS流量，而是捕获端口53的所有DNS流量。此设置仅在启用重定向dns时有效。", &cfg.CaptureAllDNS)
	flag.BindEnv(fs, constants.NetworkNamespace, "", "iptables规则应该应用的网络名称空间。", &cfg.NetworkNamespace)
	flag.BindEnv(fs, constants.CNIMode, "", "是否作为CNI插件运行。", &cfg.HostFilesystemPodNetwork)
	flag.BindEnv(fs, constants.Reconcile, "", "调和已有的和不兼容的iptables规则，而不是在检测到漂移时失败。", &cfg.Reconcile)
	flag.BindEnv(fs, constants.CleanupOnly, "", "执行强制清理，而不创建新的iptables链或规则。", &cfg.CleanupOnly)
	// This flag is a safety measure in case the idempotency changes of #50328 backfire.
	// Allow bypassing of iptables idempotency handling, and attempts to apply iptables rules regardless of table state, which may cause unrecoverable failures.
	// Consider removing it after several releases with no reported issues.
	flag.BindEnv(fs, constants.ForceApply, "", "应用iptables更改，即使它们看起来已经到位。", &cfg.ForceApply)
}

type IptablesError struct {
	Error    error
	ExitCode int
}

func GetCommand(logOpts *log.Options) *cobra.Command {
	cfg := config.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "istio-iptables",
		Short: "Set up iptables rules for Istio Sidecar",
		Long:  "istio-iptables is responsible for setting up port forwarding for Istio Sidecar.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := log.Configure(logOpts); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := cfg.FillConfigFromEnvironment(); err != nil {
				handleErrorWithCode(err, 1)
			}
			if err := cfg.Validate(); err != nil {
				handleErrorWithCode(err, 1)
			}
			if err := ProgramIptables(cfg); err != nil {
				handleErrorWithCode(err, 1)
			}

			if cfg.RunValidation {
				validator := validation_over.NewValidator(cfg)

				if err := validator.Run(); err != nil {
					// nolint: revive, stylecheck
					msg := fmt.Errorf(`iptables validation failed; workload is not ready for Istio.
When using Istio CNI, this can occur if a pod is scheduled before the node is ready.

If installed with 'cni.repair.deletePods=true', this pod should automatically be deleted and retry.
Otherwise, this pod will need to be manually removed so that it is scheduled on a node with istio-cni running, allowing iptables rules to be established.
`)
					handleErrorWithCode(msg, constants.ValidationErrorCode)
				}
			}
		},
	}
	bindCmdlineFlags(cfg, cmd)
	return cmd
}

func ProgramIptables(cfg *config.Config) error {
	ext := &dep.RealDependencies{
		HostFilesystemPodNetwork: cfg.HostFilesystemPodNetwork,
		NetworkNamespace:         cfg.NetworkNamespace,
	}

	iptConfigurator := capture_over.NewIptablesConfigurator(cfg, ext)

	if !cfg.SkipRuleApply {
		if err := iptConfigurator.Run(); err != nil {
			return err
		}
		if err := capture_over.ConfigureRoutes(cfg); err != nil {
			return fmt.Errorf("failed to configure routes: %v", err)
		}
	}
	return nil
}
