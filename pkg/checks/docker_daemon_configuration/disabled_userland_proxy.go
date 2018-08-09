package dockerdaemonconfiguration

import (
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/checks/utils"
)

type disableUserlandProxyBenchmark struct{}

func (c *disableUserlandProxyBenchmark) Definition() utils.Definition {
	return utils.Definition{
		CheckDefinition: v1.CheckDefinition{
			Name:        "CIS Docker v1.1.0 - 2.15",
			Description: "Ensure Userland Proxy is Disabled",
		}, Dependencies: []utils.Dependency{utils.InitDockerConfig},
	}
}

func (c *disableUserlandProxyBenchmark) Run() (result v1.CheckResult) {
	opts, ok := utils.DockerConfig["userland-proxy"]
	if !ok {
		utils.Warn(&result)
		utils.AddNotes(&result, "userland proxy is enabled by default")
		return
	}
	if opts.Matches("false") {
		utils.Warn(&result)
		utils.AddNotes(&result, "userland proxy is enabled")
		return
	}
	utils.Pass(&result)
	return

}

// NewDisableUserlandProxyBenchmark implements CIS-2.15
func NewDisableUserlandProxyBenchmark() utils.Check {
	return &disableUserlandProxyBenchmark{}
}
