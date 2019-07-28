package gitconfig

import (
	"errors"
	"reflect"
	"sort"
	"testing"
)

func TestShouldReturnNoErrors(t *testing.T) {
	aliases := []string{"mrs", "mr"}
	coauthorMapping := map[string]string{"mrs": "Mrs. Noujz <noujz@mrs.se>", "mr": "Mr. Noujz <noujz@mr.se>"}

	var expectedCoauthors []string
	for _, coauthor := range coauthorMapping {
		expectedCoauthors = append(expectedCoauthors, coauthor)
	}

	resolveAlias := func(alias string) (string, error) { return coauthorMapping[alias], nil }

	coauthors, errs := resolveAliases(resolveAlias)(aliases)

	if len(errs) > 0 {
		t.Errorf("unexpected errors: %s", errs)
		t.Fail()
	}

	sort.Strings(coauthors)
	sort.Strings(expectedCoauthors)

	if !reflect.DeepEqual(expectedCoauthors, coauthors) {
		t.Errorf("expected: %s, got: %s", expectedCoauthors, coauthors)
		t.Fail()
	}
}

type resolveresult struct {
	coauthor string
	err      error
}

func TestShouldAccumulateErrs(t *testing.T) {
	aliases := []string{"mrs", "mr"}
	coauthorMapping := map[string]resolveresult{"mrs": resolveresult{coauthor: "Mrs. Noujz <noujz@mrs.se>", err: nil}, "mr": resolveresult{coauthor: "", err: errors.New("failed to resolve alias mr")}}

	resolveAlias := func(alias string) (string, error) { return coauthorMapping[alias].coauthor, coauthorMapping[alias].err }

	_, errs := resolveAliases(resolveAlias)(aliases)

	if len(errs) != 1 || errs[0].Error() != "failed to resolve alias mr" {
		t.Errorf("unexpected amount of errors: %s", errs)
		t.Fail()
	}
}
