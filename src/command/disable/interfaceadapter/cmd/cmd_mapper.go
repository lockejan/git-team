package disablecmdadapter

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	commandadapter "github.com/hekmekk/git-team/src/command/adapter"
	"github.com/hekmekk/git-team/src/command/disable"
	disableeventadapter "github.com/hekmekk/git-team/src/command/disable/interfaceadapter/event"
	statuscmdmapper "github.com/hekmekk/git-team/src/command/status/interfaceadapter/cmd"
	activation "github.com/hekmekk/git-team/src/shared/activation/impl"
	configds "github.com/hekmekk/git-team/src/shared/config/datasource"
	gitconfig "github.com/hekmekk/git-team/src/shared/gitconfig/impl"
	state "github.com/hekmekk/git-team/src/shared/state/impl"
)

// Command the disable command
func Command(root commandadapter.CommandRoot) *kingpin.CmdClause {
	disable := root.Command("disable", "Use default commit template and remove prepare-commit-msg hook")

	disable.Action(commandadapter.Run(policy(), disableeventadapter.MapEventToEffectsFactory(statuscmdmapper.Policy())))

	return disable
}

func policy() disable.Policy {
	return disable.Policy{
		Deps: disable.Dependencies{
			ConfigReader:        configds.NewGitconfigDataSource(gitconfig.NewDataSource()),
			GitConfigReader:     gitconfig.NewDataSource(),
			GitConfigWriter:     gitconfig.NewDataSink(),
			StatFile:            os.Stat,
			RemoveFile:          os.RemoveAll,
			StateWriter:         state.NewGitConfigDataSink(gitconfig.NewDataSink()),
			ActivationValidator: activation.NewGitConfigDataSource(gitconfig.NewDataSource()),
		},
	}
}
