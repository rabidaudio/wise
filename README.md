# Wise

**Note:** work in progress. Thoughts and contributions welcome

Wise is a command line tool for version control which is easier to learn and harder to mess up than `git`, while still leveraging the power of `git` under the hood, making it compatible with existing repositories.

Linus [named it "git" after the British slang for "idiot"](https://en.wikipedia.org/wiki/Git#Naming) because [`git` is stupid](https://www.youtube.com/watch?v=4XpnKHJAok8). It was built 16 years ago for managing the linux kernel repository and a game changer in the industry. It blew the other version control systems out of the water because the creators learned from many of the mistakes of earlier version control tools.

However software development processes and `git`'s usage has evolved over time and some of it's conventions and names for things have lost their original meaning. Many of the [top Stack Overflow questions of all time](https://stackoverflow.com/questions?sort=votes) are about how to do basic things in `git`. It's hard enough if you're new to programming to understand how version control works, you shouldn't have to also learn the quirks of poor vocabulary and memorize a bunch of command line flags.

Unlike `git`, Wise is wise. The goal is to make an easy to learn and use command line tool with a very limited surface area targeted to the most common use cases and mistakes. Since it's built on Git, it's fully compatible with your existing code base and repository management service (GitHub, GitLab, BitBucket, etc), and you can always fall back to using `git` when you need more control. For this reason, we are going to be judicious about keeping the functionality small and understandable, while also being aggressive at detecting and interrupting common mistakes.

## Common "How do I X in Git?" questions

- undo commits
- delete branches
- rename branches
- overwrite local changes with a pull
- modify an un-pushed commit

## Confusing things about Git:

- `checkout` creates a branch, but `branch` deletes it
- `clone`, `fetch`, `pull`, and `checkout` all have similar english meanings but very different `git` behaviors
- should I `merge` or `rebase`?
- staging changes before committing
- stashing to switch branches real quick with the intention of returning

## Clunky user experience things:

- even for git experts, it's easy to use the wrong name + email address. often your email address will be different for different projects (e.g. work and personal)
- `git` wasn't designed specifically for a single remote source-of-truth and a git flow pattern even though this is the most common setup by far
- ignoring files is a huge pain, and a security risk if done wrong. removing files that hadn't been ignored is not intuitive
- `GIT_EDITOR` defaults to `vim` which is a disaster for people unfamiliar with `vim`. `nano` is not bad, but there's really no reason you need a whole terminal text editor to write a commit message

## Roadmap

- [ ] Basic git flow features, with a condensed, clear vocabulary
- [ ] Solutions to the most common stack overflow questions
- [ ] Better project setup (cloning and intializing)
- [ ] Automation around `.gitignore` perhaps via https://github.com/github/gitignore with detection of likely gitignore mistakes and remediation
- [ ] Automation around setting username/email per-project


# Notes:

```
// wi status -- print anything useful for understanding what's happening. cwd, branch,
//			does a remote branch exist, what are the differences in commits between local and remote,
//			are there uncommitted changes
// wi branch -- create a new branch. confirm if source branch is not a default/protected one. fetch changes from remote before branching
// wi commit -- look for unresolved merge conflicts and error,
// 			interactively stage and unstage files, encourage review, supply a message, and commit
// wi push -- pull down changes from source branch and merge, then push to remote, creating a new remote branch if necessary. error if
//			pushing to a protected branch
// wi sync? -- fetch, check for changes to main branch and merge/rebase. if there are existing changes, offer to apply them on top or throw
//			them away
// wi peek <ref> -- stash any current changes and check out a particular commit/branch/tag. fetch it from remote if required.
//			block committing until switched back with `wi return`?
// wi undo-commit -- check if commit has been pushed already, then revert last commit and unstage changes
// wi switch? -- stash any current changes on the existing branch and switch to another existing branch. difference from `peek`
//			being you can commit
// wi learn -- interactive tutorial for version control and git flow

// need consistency - noun-oriented or verb-oriented?
// "branch" and "commit" are each both nouns and verbs
// powerful terminal uis: https://github.com/rivo/tview
```
