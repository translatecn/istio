package main

import (
	"regexp"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type vmContext struct {
	// 嵌入默认的 VM 上下文，这样我们就不需要重新实现所有方法
	types.DefaultVMContext
}

func (ctx *vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// 嵌入默认的插件上下文，这样我们就不需要重新实现所有方法
	types.DefaultPluginContext

	pattern     string
	replaceWith string
	configData  string // 保存插件的一些配置信息
}

// 注入额外的 Header
var additionalHeaders = map[string]string{
	"who-am-i":    "go-wasm-demo",
	"injected-by": "istio-api!",
	"site":        "youdianzhishi.com",
	"author":      "阳明",
	// 定义自定义的header，每个返回中都添加以上header
}

// NewHttpContext 为每个 HTTP 请求创建一个新的上下文。
func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpRegex{
		contextID:     contextID,
		pluginContext: ctx,
	}
}

// OnPluginStart 在插件被加载时调用。
func (ctx *pluginContext) OnPluginStart(pluginCfgSize int) types.OnPluginStartStatus {
	proxywasm.LogWarnf("regex/main.go OnPluginStart()")
	// 获取插件配置
	data, err := proxywasm.GetPluginConfiguration()
	if data == nil {
		return types.OnPluginStartStatusOK
	}
	if err != nil {
		proxywasm.LogWarnf("failed read plug-in config: %v", err)
		return types.OnPluginStartStatusFailed
	}

	proxywasm.LogWarnf("read plug-in config: %s\n", string(data))

	// 插件启动的时候读取配置
	ctx.configData = string(data)
	ctx.pattern = "banana/([0-9]*)"
	ctx.replaceWith = "status/$1"

	return types.OnPluginStartStatusOK
}

// OnPluginDone 在插件被卸载时调用。
func (ctx *pluginContext) OnPluginDone() bool {
	proxywasm.LogWarnf("regex/main.go OnPluginDone()")
	return true
}

type httpRegex struct {
	// 嵌入默认的 HTTP 上下文，这样我们就不需要重新实现所有方法
	types.DefaultHttpContext
	// contextID 是插件上下文的 ID，它是唯一的。
	contextID     uint32
	pluginContext *pluginContext
}

// OnHttpResponseHeaders 在收到 HTTP 响应头时调用。
func (ctx *httpRegex) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogWarnf("%d httpRegex.OnHttpResponseHeaders(%d, %t)", ctx.contextID, numHeaders, endOfStream)

	// 添加 Header
	for k, v := range additionalHeaders {
		if err := proxywasm.AddHttpResponseHeader(k, v); err != nil {
			proxywasm.LogWarnf("failed to add response header %s: %v", k, err)
		}
	}

	//为了便于演示观察，将配置信息也加到返回头里
	proxywasm.AddHttpResponseHeader("configData", ctx.pluginContext.configData)
	return types.ActionContinue
}

// OnHttpRequestHeaders 在收到 HTTP 请求头时调用。
func (ctx *httpRegex) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogWarnf("%d httpRegex.OnHttpRequestHeaders(%d, %t)", ctx.contextID, numHeaders, endOfStream)

	re := regexp.MustCompile(ctx.pluginContext.pattern)
	replaceWith := ctx.pluginContext.replaceWith

	s, err := proxywasm.GetHttpRequestHeader(":path")
	if err != nil {
		proxywasm.LogWarnf("Could not get request header: %v", err)
	} else {
		result := re.ReplaceAllString(s, replaceWith)
		proxywasm.LogWarnf("path: %s, result: %s", s, result)

		err = proxywasm.ReplaceHttpRequestHeader(":path", result)
		if err != nil {
			proxywasm.LogWarnf("Could not set request header to %q: %v", result, err)
		}
	}

	return types.ActionContinue
}

func (ctx *httpRegex) OnHttpStreamDone() {
	proxywasm.LogWarnf("%d OnHttpStreamDone", ctx.contextID)
}

func main() {
	proxywasm.LogWarnf("regex/main.go main() REACHED")
	// 设置 VM 上下文，这样我们就可以在插件启动时读取配置。
	proxywasm.SetVMContext(&vmContext{})
}
