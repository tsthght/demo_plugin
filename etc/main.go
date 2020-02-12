package main

import (
	"fmt"
	"github.com/tsthght/demo_plugin/plugin"
)

func main() {
	// 插件管理链表
	pluginList := make(map[string]plugin.MyPlugin)

	plugin_name := "demo.so"
	plugin_path := "bin"
	// 根据路径/名字加载插件
	e, p := plugin.LoadPlugin(plugin_path, plugin_name)
	if e != nil {
		fmt.Printf(e.Error())
		return
	}
	pluginList[plugin_name] = p
	//插件的调用
	for _, v := range pluginList {
		v.DoFilter()
	}
}
