`git just`: A more intuitive CLI subcommand for git
---
# Vision
git is a beautiful version control system, 
and an incredibly powerful collaborative
tool for working in teams. However, the
suite of command line tools that come 
out-of-the-box are not great. `git just`
aims to make it simpler to interact with
git repositories for common tasks, for 
new users and experts alike.

# Simplifications
`git just` will endeavor to maintain the following invariants, to reduce the amount of mental load on the user.
These simplifications are designed to not reduce the expressive power of the git data model, and
are usable by git wizards as well as novices.

- `origin/branch` should always point to the _latest thing actually available on the server_, 
        as long as you have internet access.
- `origin/branch` and `branch` should always point to the same thing, 
        and it shall be referred to as `branch` 
- `add` and `commit` shall no longer be two separate operations. There shall only be one set of 
        uncommited files, and they shall be the version in your filesystem.
- Editing history that has not been pushed should be easy (amending, squashing, etc.) but editing 
    history that has already been pushed shall always make you feel bad.
- If your project lives on Github, Gitlab, or another common host, we should take advantage of that.

# Progress
- [] `git just download`: Downloads a git repository.
    - [] Does the right thing with URLs from Github (or others)
    - [] Prefers SSH URL's when the machine is authenticated to use them properly
        - [] Offers to help setup SSH keys in a reasonable way if not currently configured
    - [] Uses Github account information to automatically resolve unqualified repo names
- [] `git just look`: Basically `git status`, probably with some tweaks
    - [] Runs a `git fetch` and updates all inactive branches
    - Maybe interactive?
- [] `git just save`: Commits changes since you last saved, and pushes them
    - Figure out whether this should actually push
    - [] Bring up nice in-terminal UI for creating a commit message 
        - [] ...and selecting which files to save
- [] `git just merge`: Provides options for useful merges to make
    - [] Provides a preview of the result
    - [] Includes or launches a merge tool
    - [] Can make and interact with pull requests in Github (or others) 
- [] `git just cd`: From anywhere on your system, changes into the specified repo 
                    (if `git-just` has ever run there)
- [] `git just configure`: Enables repository and system configuration
    - [] Common git system settings (username, email, push.matching.default)
    - [] git-just repository defaults (personal project? Trunk development? Pull requests?)
    - [] git-just system settings (GitHub login, etc...)
    - [] system aliases for git-just and subcommands
    - [] SSH keys and machine settings that matter for using git
