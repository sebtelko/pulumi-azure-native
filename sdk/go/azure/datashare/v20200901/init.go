// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:datashare/v20200901:Account":
		r = &Account{}
	case "azure-native:datashare/v20200901:DataSet":
		r = &DataSet{}
	case "azure-native:datashare/v20200901:DataSetMapping":
		r = &DataSetMapping{}
	case "azure-native:datashare/v20200901:Invitation":
		r = &Invitation{}
	case "azure-native:datashare/v20200901:Share":
		r = &Share{}
	case "azure-native:datashare/v20200901:ShareSubscription":
		r = &ShareSubscription{}
	case "azure-native:datashare/v20200901:SynchronizationSetting":
		r = &SynchronizationSetting{}
	case "azure-native:datashare/v20200901:Trigger":
		r = &Trigger{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"datashare/v20200901",
		&module{version},
	)
}
