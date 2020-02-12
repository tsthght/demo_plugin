package main

import (
	"fmt"

	"github.com/tsthght/demo_plugin/plugin"
)

type PluginDemo struct {}
func (pd PluginDemo) DoFilter() {
	fmt.Printf("I am PluginDemo\n")
}

func NewPlugin() plugin.MyPlugin {
	return PluginDemo{}
}

