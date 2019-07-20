package status

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/hekmekk/git-team/src/config"
)

func PersistEnabled(coauthors []string) error {
	return persist(state{Status: enabled, Coauthors: coauthors})
}

func PersistDisabled() error {
	return persist(state{Status: disabled})
}

func Print() {
	status := fetch()
	fmt.Println(status.toString())
}

func persist(state state) error {
	deps := persistDependencies{
		loadConfig: config.Load,
		writeFile:  ioutil.WriteFile,
	}
	return persistFactory(deps)(state)
}

func fetch() state {
	deps := loadDependencies{
		loadConfig:     config.Load,
		tomlDecodeFile: toml.DecodeFile,
	}
	return loadStatusFromFileFactory(deps)()
}
