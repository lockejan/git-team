package disablecmdadapter

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	commandadapter "github.com/hekmekk/git-team/src/command/adapter"
	"github.com/hekmekk/git-team/src/command/disable"
	disableeventadapter "github.com/hekmekk/git-team/src/command/disable/interfaceadapter/event"
	statuscmdmapper "github.com/hekmekk/git-team/src/command/status/interfaceadapter/cmd"
	staterepository "github.com/hekmekk/git-team/src/core/state_repository"
	gitconfig "github.com/hekmekk/git-team/src/shared/gitconfig/impl"
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
			GitConfigReader: gitconfig.NewDataSource(),
			GitConfigWriter: gitconfig.NewDataSink(),
			StatFile:        os.Stat,
			RemoveFile:      os.Remove,
			PersistDisabled: staterepository.PersistDisabled,
		},
	}
}
