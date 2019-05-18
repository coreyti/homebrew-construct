package add

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/coreyti/construct/models"
)

type Runner struct {
	path string
	repo string
}

func NewRunner(path string, repo string) *Runner {
	return &Runner{
		path: path,
		repo: repo,
	}
}

func (r *Runner) Run() error {
	repo := models.NewRepository(r.repo)
	if err := repo.Sync(r.path); err != nil {
		fmt.Printf("Syncing repo; error: %s", err)
		os.Exit(1)
	}

	// TODO: let errors raise
	return nil
}

func (r *Runner) TmpInstall() {
	bundle := strings.Join([]string{r.path, "Brewfile"}, "/")

	if _, err := os.Stat(bundle); os.IsNotExist(err) {
		fmt.Printf("Skipping `brew bundle` for %s", r.repo)
	} else {
		fileArg := strings.Join([]string{"--file", bundle}, "=")
		arguments := []string{"bundle", "-v", fileArg}
		execution := exec.Command("brew", arguments...)

		_, err := execution.CombinedOutput()
		if err != nil {
			fmt.Printf("Homebrew bundling; error: %s", err)
			os.Exit(1)
		}
	}
}
