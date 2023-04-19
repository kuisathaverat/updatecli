package tag

import (
	"fmt"

	"github.com/updatecli/updatecli/pkg/core/pipeline/scm"
)

// Target ensure that a specific release exist on Gitlab, otherwise creates it
func (g Gitlab) Target(source string, scm scm.ScmHandler, dryRun bool) (bool, []string, string, error) {
	return false, []string{}, "", fmt.Errorf("target not supported for the plugin Gitlab Tags")
}
