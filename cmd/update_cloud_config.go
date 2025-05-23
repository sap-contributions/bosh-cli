package cmd

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	. "github.com/cloudfoundry/bosh-cli/v7/cmd/opts" //nolint:staticcheck
	boshdir "github.com/cloudfoundry/bosh-cli/v7/director"
	boshtpl "github.com/cloudfoundry/bosh-cli/v7/director/template"
	boshui "github.com/cloudfoundry/bosh-cli/v7/ui"
)

type UpdateCloudConfigCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewUpdateCloudConfigCmd(ui boshui.UI, director boshdir.Director) UpdateCloudConfigCmd {
	return UpdateCloudConfigCmd{ui: ui, director: director}
}

func (c UpdateCloudConfigCmd) Run(opts UpdateCloudConfigOpts) error {
	tpl := boshtpl.NewTemplate(opts.Args.CloudConfig.Bytes)

	bytes, err := tpl.Evaluate(opts.VarFlags.AsVariables(), opts.OpsFlags.AsOp(), boshtpl.EvaluateOpts{}) //nolint:staticcheck
	if err != nil {
		return bosherr.WrapErrorf(err, "Evaluating cloud config")
	}

	cloudConfigDiff, err := c.director.DiffCloudConfig(opts.Name, bytes)
	if err != nil {
		return err
	}

	diff := NewDiff(cloudConfigDiff.Diff)
	diff.Print(c.ui)

	err = c.ui.AskForConfirmation()
	if err != nil {
		return err
	}

	return c.director.UpdateCloudConfig(opts.Name, bytes)
}
