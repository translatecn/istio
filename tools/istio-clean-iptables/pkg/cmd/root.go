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
	"github.com/spf13/cobra"

	"istio.io/istio/pkg/flag"
	"istio.io/istio/pkg/log"
	"istio.io/istio/tools/istio-clean-iptables/pkg/config"
	"istio.io/istio/tools/istio-iptables/pkg/constants"
)

func bindCmdlineFlags(cfg *config.Config, cmd *cobra.Command) {
	fs := cmd.Flags()
	flag.BindEnv(fs, constants.DryRun, "n", "不要调用任何外部依赖，比如iptables。", &cfg.DryRun)
	flag.BindEnv(fs, constants.ProxyUID, "u", "指定不应用重定向的用户的UID。通常，这是代理容器的UID。", &cfg.ProxyUID)
	flag.BindEnv(fs, constants.ProxyGID, "g", "指定不应用重定向的用户的GID（与-u param的默认值相同）。", &cfg.ProxyGID)
	flag.BindEnv(fs, constants.RedirectDNS, "", "启用istio-agent捕获dns流量功能。", &cfg.RedirectDNS)
	// Allow binding to a different var, for consistency with other components
	flag.AdditionalEnv(fs, constants.RedirectDNS, "ISTIO_META_DNS_CAPTURE")

	flag.BindEnv(fs, constants.CaptureAllDNS, "", "不是只捕获到DNS服务器IP的DNS流量，而是捕获端口53的所有DNS流量。此设置仅在启用重定向dns时有效。", &cfg.CaptureAllDNS)
	flag.BindEnv(fs, constants.InboundInterceptionMode, "m", "用于将入站连接重定向到Envoy的模式，可以是“redirect”或“TPROXY”。", &cfg.InboundInterceptionMode)
	flag.BindEnv(fs, constants.InboundTProxyMark, "t", "", &cfg.InboundTProxyMark)
}

func GetCommand(logOpts *log.Options) *cobra.Command {
	cfg := config.DefaultConfig()
	cmd := &cobra.Command{
		Use:   "istio-clean-iptables",
		Short: "Clean up iptables rules for Istio Sidecar",
		Long:  "Script responsible for cleaning up iptables rules",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := log.Configure(logOpts); err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.FillConfigFromEnvironment()
			if err := cfg.Validate(); err != nil {
				return err
			}
			ext := NewDependencies(cfg)

			iptVer, err := ext.DetectIptablesVersion(false)
			if err != nil {
				return err
			}
			ipt6Ver, err := ext.DetectIptablesVersion(true)
			if err != nil {
				return err
			}

			cleaner := NewIptablesCleaner(cfg, &iptVer, &ipt6Ver, ext)
			cleaner.Run()
			return nil
		},
	}
	bindCmdlineFlags(cfg, cmd)
	return cmd
}
