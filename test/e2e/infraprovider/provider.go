package infraprovider

import (
	"github.com/ovn-org/ovn-kubernetes/test/e2e/infraprovider/api"
	"github.com/ovn-org/ovn-kubernetes/test/e2e/infraprovider/providers/openshift"

	"k8s.io/client-go/rest"
)

type Name string

func (n Name) String() string {
	return string(n)
}

var infraProvider api.Provider

// Set infrastructure provider.
func Set(provider api.Provider) {
	infraProvider = provider
}

// SetAuto detects which infrastructure provider. Arg config is not needed for KinD provider but downstream implementations
// will require access to the kapi to infer what platform k8 is running on.
func SetAuto(config *rest.Config) error {
	infraProvider = openshift.NewOpenShiftProvider(config)
	return nil
}

// Get infrastructure provider.
func Get() api.Provider {
	if infraProvider == nil {
		panic("infra provider not set")
	}
	return infraProvider
}
