package plugin

import (
	"plugin"
	"fmt"
	"errors"
)

const FactorFunc = "NewPlugin"

func LoadPlugin(plugin_path, plugin_name string) (error, MyPlugin) {
	fp := plugin_path + "/" + plugin_name
	p, err := plugin.Open(fp)
	if err != nil {
		return errors.New(fmt.Sprintf("Open %s failed, err: %s\n", fp, err.Error())), nil
	}
	f, err := p.Lookup(FactorFunc)
	if err != nil {
		return errors.New(fmt.Sprintf("Lookup %s faled, err: %s\n", FactorFunc, err.Error())), nil
	}

	newPlugin, ok := f.(func() MyPlugin)
	if !ok {
		return errors.New(fmt.Sprintf("Type of %T is not match\n", f)), nil
	}
	return nil, newPlugin()
}
