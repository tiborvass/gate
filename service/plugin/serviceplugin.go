// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugin

import (
	"fmt"
	"reflect"

	"github.com/tsavola/gate/plugin"
	"github.com/tsavola/gate/service"
)

// Function names exported by service plugins.
const (
	ServiceConfigSymbol = "ServiceConfig" // Optional func() interface{}
	InitServicesSymbol  = "InitServices"  // Required func(service.Config) error
)

type ServicePlugins struct {
	plugin.Plugins
	ServiceConfig map[string]interface{}
}

// OpenAll generic and service plugins found under libdir.
func OpenAll(libdir string) (result ServicePlugins, err error) {
	result.Plugins, err = plugin.OpenAll(libdir)
	if err != nil {
		return
	}

	result.ServiceConfig = make(map[string]interface{})

	for _, p := range result.Plugins.Plugins {
		var x interface{}

		x, err = getServiceConfig(p)
		if err != nil {
			return
		}

		if x != nil {
			result.ServiceConfig[p.Name] = x
		}
	}

	return
}

func getServiceConfig(p plugin.Plugin) (interface{}, error) {
	x, err := p.Lookup(ServiceConfigSymbol)
	if err != nil {
		return nil, nil
	}

	f, ok := x.(func() interface{})
	if !ok {
		return nil, fmt.Errorf("%s: %s is a %s; expected a %s", p, ServiceConfigSymbol, reflect.TypeOf(x), reflect.TypeOf(f))
	}

	return f(), nil
}

func (ps ServicePlugins) InitServices(config service.Config) (err error) {
	for _, p := range ps.Plugins.Plugins {
		_, hasConfig := ps.ServiceConfig[p.Name]

		err = initServices(config, p, hasConfig)
		if err != nil {
			return
		}
	}

	return
}

func initServices(config service.Config, p plugin.Plugin, require bool) error {
	x, err := p.Lookup(InitServicesSymbol)
	if err != nil {
		if require {
			return fmt.Errorf("%s: %v", p, err)
		} else {
			return nil
		}
	}

	f, ok := x.(func(service.Config) error)
	if !ok {
		return fmt.Errorf("%s: %s is a %s; expected a %s", p, InitServicesSymbol, reflect.TypeOf(x), reflect.TypeOf(f))
	}

	return f(config)
}
