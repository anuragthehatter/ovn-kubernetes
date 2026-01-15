package deploymentconfig

import (
	"github.com/ovn-org/ovn-kubernetes/test/e2e/deploymentconfig/api"
	"github.com/ovn-org/ovn-kubernetes/test/e2e/deploymentconfig/configs/openshift"
)

var deploymentConfig api.DeploymentConfig

// Set deployment config.
func Set(deployment api.DeploymentConfig) {
	deploymentConfig = deployment
}

// SetAuto auto-detects and sets the deployment config.
func SetAuto() {
	// upstream currently uses KinD as its preferred platform infra, so if we detect KinD, its upstream
	if openshift.IsBaremetalds() {
		deploymentConfig = openshift.NewBaremetalds()
	}
	if deploymentConfig == nil {
		panic("failed to determine the deployment config")
	}
}

// Get deployment config.
func Get() api.DeploymentConfig {
	if deploymentConfig == nil {
		panic("deployment config type not set")
	}
	return deploymentConfig
}
