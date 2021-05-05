package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "wise",
		Description: "Smarter version control than git",
		Authors: []*cli.Author{
			{
				Name:  "Charles Julian Knight",
				Email: "julian@fixdapp.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "status",
				Aliases:     []string{"s"},
				Action:      withRepo(status),
				Description: "Check the current status of the working tree",
			},
		},
		// Action: withRepo(status),
	}
	err := app.Run(os.Args)
	if err == nil {
		return
	}
	if _, ok := err.(FriendlyError); ok {
		fmt.Println("It's unlikely this is an error with Wise.")
	} else {
		// TODO: improve experience
		fmt.Println("It's likely this is an error with Wise itself. Consider reporting it by opening an issue.")
		panic(err)
	}
}

func withRepo(handler func(ctx *cli.Context, repo *git.Repository, wt *git.Worktree) error) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("unable to determine current directory: %w", err)
		}
		repo, err := git.PlainOpenWithOptions(dir, &git.PlainOpenOptions{
			DetectDotGit: true,
		})
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				return notARepositoryError(dir)
			}
			return fmt.Errorf("get repo: %w", err)
		}
		wt, err := repo.Worktree()
		if err != nil {
			// if errors.Is(err, git.ErrIsBareRepository) {
			// 	// TODO: handle this
			// }
			return fmt.Errorf("get worktree: %w", err)
		}
		return handler(ctx, repo, wt)
	}
}

func notARepositoryError(cwd string) error {
	// TODO: improve directions on how to resolve this issue
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "The current directory `%s` does not appear to be in a git repository.\n\n", cwd)
	fmt.Fprint(buf, "It could be that you aren't in the directory of your project, or it could be that you want to create a new repository.")
	return NewFriendlyError(buf.String())
}

func status(ctx *cli.Context, repo *git.Repository, wt *git.Worktree) error {
	status, err := wt.Status()
	if err != nil {
		return fmt.Errorf("get status: %w", err)
	}

	// TODO: print meaningful information about the current state
	fmt.Println(status.String())

	return nil
}
