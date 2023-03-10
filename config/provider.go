/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	// "github.com/upbound/upjet-provider-template/config/null"
	"github.com/devops-chris/provider-twingate/config/connector"
	"github.com/devops-chris/provider-twingate/config/remote_network"
)

const (
	resourcePrefix = "template"
	modulePath     = "github.com/upbound/upjet-provider-template"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		// null.Configure,
		remote_network.Configure,
		connector.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
