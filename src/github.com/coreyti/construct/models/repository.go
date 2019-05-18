package models

import (
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

type Repository struct {
	uri string
}

func NewRepository(uri string) *Repository {
	return &Repository{
		uri: uri,
	}
}

func (r *Repository) Sync(path string) error {
	var repo *git.Repository

	if _, err := os.Stat(path); os.IsNotExist(err) {
		repo, err = git.PlainClone(path, false, &git.CloneOptions{
			URL: r.uri,
		})

		if err != nil {
			fmt.Printf("Cloning repo; error: %s", err)
			os.Exit(1)
		}
	} else {
		repo, err = git.PlainOpen(path)
		if err != nil {
			fmt.Printf("Opening repo; error: %s", err)
			os.Exit(1)
		}
	}

	work, err := repo.Worktree()
	if err != nil {
		fmt.Printf("Getting repo work tree; error: %s", err)
		os.Exit(1)
	}

	if err := work.Checkout(&git.CheckoutOptions{}); err != nil {
		return fmt.Errorf("Checking out repo; error: %s", err)
	}

	// TODO: let errors raise
	return nil
}
