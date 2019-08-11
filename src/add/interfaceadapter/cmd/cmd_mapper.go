package addcmdadapter

import (
	"bufio"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hekmekk/git-team/src/add"
	"github.com/hekmekk/git-team/src/gitconfig"
	"github.com/hekmekk/git-team/src/validation"
)

// Definition the command, arguments, and dependencies
type Definition struct {
	CommandName string
	Request     add.AssignmentRequest
	Deps        add.Dependencies
}

// New the constructor for Definition
func New(app *kingpin.Application) Definition {

	command := app.Command("add", "Add a new or override an existing alias to co-author assignment")
	alias := command.Arg("alias", "The alias to assign a co-author to").Required().String()
	coauthor := command.Arg("coauthor", "The co-author").Required().String()

	return Definition{
		CommandName: command.FullCommand(),
		Request: add.AssignmentRequest{
			Alias:    alias,
			Coauthor: coauthor,
		},
		Deps: add.Dependencies{
			SanityCheckCoauthor: validation.SanityCheckCoauthor,
			GitAddAlias:         gitconfig.AddAlias,
			GitResolveAlias:     gitconfig.ResolveAlias,
			StdinReadLine:       func() (string, error) { return bufio.NewReader(os.Stdin).ReadString('\n') },
		},
	}
}