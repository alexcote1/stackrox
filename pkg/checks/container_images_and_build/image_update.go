package containerimagesandbuild

import (
	"strings"

	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/checks/utils"
	"github.com/stackrox/rox/pkg/docker"
)

type imageUpdateInstructionsBenchmark struct{}

var updateCmds = []string{
	"apk update",
	"apt update",
	"apt-get update",
	"yum update",
}

func (c *imageUpdateInstructionsBenchmark) Definition() utils.Definition {
	return utils.Definition{
		CheckDefinition: v1.CheckDefinition{
			Name:        "CIS Docker v1.1.0 - 4.7",
			Description: "Ensure update instructions are not use alone in the Dockerfile",
		}, Dependencies: []utils.Dependency{utils.InitImages},
	}
}

func (c *imageUpdateInstructionsBenchmark) Run() (result v1.CheckResult) {
	utils.Pass(&result)
	for _, image := range utils.Images {
		ctx, cancel := docker.TimeoutContext()
		defer cancel()
		historySlice, err := utils.DockerClient.ImageHistory(ctx, image.ID)
		if err != nil {
			utils.Warn(&result)
			utils.AddNotef(&result, "Could not get image history for image %v: %+v", utils.GetReadableImageName(image), err)
			continue
		}
		for _, history := range historySlice {
			cmd := strings.ToLower(history.CreatedBy)
			cmd = strings.Replace(cmd, "\t", "", -1)
			cmd = strings.TrimPrefix(cmd, "/bin/sh -c #(nop)")
			cmd = strings.TrimPrefix(cmd, "/bin/sh -c")
			for _, updateCmd := range updateCmds {
				if cmd == updateCmd {
					utils.Warn(&result)
					utils.AddNotef(&result, "Image '%v' has an update command alone in layer: '%v'", utils.GetReadableImageName(image), cmd)
				}
			}
		}
	}
	return
}

// NewImageUpdateInstructionsBenchmark implements CIS-4.7
func NewImageUpdateInstructionsBenchmark() utils.Check {
	return &imageUpdateInstructionsBenchmark{}
}
