---
name: gh-cli
description: Interacts with GitHub via the gh CLI for managing pull requests, issues, repos, releases, workflows, gists, and the GitHub API. Use when the user needs to work with GitHub resources from the command line.
allowed-tools: Bash(gh-cli:*)
---

# GitHub CLI (gh)

## Authentication

```bash
# Login interactively (browser-based)
gh auth login
# Login with a token
gh auth login --with-token < token.txt
# Check auth status
gh auth status
# Switch between accounts
gh auth switch
# Get current token
gh auth token
```

## Commands

### Pull Requests

```bash
# Create
gh pr create --title "Title" --body "Description"
gh pr create --fill                          # auto-fill from commits
gh pr create --fill-verbose                  # include commit bodies
gh pr create --draft                         # create as draft
gh pr create --base develop --head feature   # specify branches
gh pr create --reviewer user1,user2          # request reviewers
gh pr create --label bug --label urgent      # add labels
gh pr create --assignee @me                  # self-assign
gh pr create --web                           # open in browser
gh pr create --dry-run                       # preview without creating

# List
gh pr list
gh pr list --author "@me"
gh pr list --state merged
gh pr list --label bug --label "priority 1"
gh pr list --search "status:success review:required"
gh pr list --json number,title,state         # JSON output
gh pr list --limit 100

# View
gh pr view 123
gh pr view 123 --json title,body,state
gh pr view 123 --comments
gh pr view --web                             # current branch PR in browser

# Checkout
gh pr checkout 123
gh pr checkout 123 --force                   # reset local branch
gh pr checkout --detach

# Diff
gh pr diff 123
gh pr diff --name-only                       # changed file names only
gh pr diff --patch                           # patch format

# Checks
gh pr checks 123
gh pr checks --watch                         # monitor until complete
gh pr checks --watch --fail-fast             # exit on first failure
gh pr checks --required                      # only required checks
gh pr checks --json name,state,bucket

# Review
gh pr review --approve
gh pr review --comment -b "Looks good"
gh pr review 123 --request-changes -b "Please fix X"
gh pr review 123 --approve -b "LGTM"

# Merge
gh pr merge 123
gh pr merge --squash                         # squash merge
gh pr merge --rebase                         # rebase merge
gh pr merge --merge                          # merge commit
gh pr merge --auto                           # auto-merge when ready
gh pr merge --delete-branch                  # delete branch after
gh pr merge --squash --subject "feat: thing" --body "Details"
gh pr merge --admin                          # bypass checks

# Edit
gh pr edit 123 --title "New title"
gh pr edit 123 --add-label bug
gh pr edit 123 --add-reviewer user1
gh pr edit 123 --base main

# Other
gh pr close 123
gh pr reopen 123
gh pr ready 123                              # mark draft as ready
gh pr lock 123
gh pr unlock 123
gh pr revert 123
gh pr update-branch 123
gh pr comment 123 --body "Comment text"
gh pr status                                 # PRs relevant to you
```

### Issues

```bash
# Create
gh issue create --title "Bug" --body "Description"
gh issue create --label bug --label "help wanted"
gh issue create --assignee @me
gh issue create --milestone "v1.0"
gh issue create --project "Roadmap"
gh issue create --template "Bug Report"
gh issue create --web

# List
gh issue list
gh issue list --assignee "@me"
gh issue list --author monalisa
gh issue list --label bug
gh issue list --state closed
gh issue list --state all
gh issue list --milestone "The big 1.0"
gh issue list --search "error no:assignee sort:created-asc"
gh issue list --json number,title,state,labels

# View
gh issue view 123
gh issue view 123 --comments
gh issue view 123 --json title,body,state
gh issue view 123 --web

# Edit
gh issue edit 123 --title "Updated title"
gh issue edit 123 --add-label bug
gh issue edit 123 --add-assignee user1

# Other
gh issue close 123
gh issue reopen 123
gh issue delete 123
gh issue pin 123
gh issue unpin 123
gh issue lock 123
gh issue unlock 123
gh issue transfer 123 target/repo
gh issue comment 123 --body "Comment text"
gh issue develop 123                         # create branch from issue
gh issue status                              # issues relevant to you
```

### Repositories

```bash
# Clone
gh repo clone owner/repo
gh repo clone owner/repo -- --depth=1        # shallow clone
gh repo clone owner/repo target-dir

# Create
gh repo create my-repo --public
gh repo create my-repo --private
gh repo create my-repo --public --clone      # create and clone
gh repo create --source=. --remote=origin --push  # from existing local

# List
gh repo list
gh repo list owner --limit 100
gh repo list --language go

# View
gh repo view
gh repo view owner/repo
gh repo view --web
gh repo view --json name,description,url

# Other
gh repo fork owner/repo
gh repo fork owner/repo --clone
gh repo edit --default-branch main
gh repo edit --visibility private
gh repo sync                                 # sync fork with upstream
gh repo archive owner/repo
gh repo unarchive owner/repo
gh repo delete owner/repo --yes
gh repo rename new-name
gh repo set-default owner/repo
```

### Releases

```bash
# Create
gh release create v1.0.0
gh release create v1.0.0 --title "Release v1.0.0" --notes "Release notes"
gh release create v1.0.0 -F changelog.md    # notes from file
gh release create v1.0.0 --generate-notes   # auto-generate notes
gh release create v1.0.0 --draft
gh release create v1.0.0 --prerelease
gh release create v1.0.0 --target main
gh release create v1.0.0 ./dist/*.tar.gz    # upload assets
gh release create v1.0.0 './dist/app.zip#My App'  # asset with label
gh release create v1.0.0 --notes-from-tag
gh release create v1.0.0 --discussion-category "Announcements"

# List / View
gh release list
gh release view v1.0.0
gh release view --web

# Download
gh release download v1.0.0
gh release download v1.0.0 --pattern "*.tar.gz"
gh release download --latest

# Other
gh release edit v1.0.0 --draft=false
gh release upload v1.0.0 ./dist/*.tar.gz
gh release delete v1.0.0 --yes
```

### Workflow Runs (GitHub Actions)

```bash
# List runs
gh run list
gh run list --workflow ci.yml
gh run list --branch main
gh run list --status failure
gh run list --user monalisa
gh run list --json databaseId,status,conclusion,name

# View run
gh run view 12345
gh run view 12345 --verbose                  # show job steps
gh run view 12345 --log                      # full log output
gh run view 12345 --log-failed               # logs for failed steps
gh run view 12345 --job 456789               # specific job
gh run view 12345 --web
gh run view 12345 --exit-status              # non-zero if failed
gh run view 12345 --attempt 3               # specific attempt

# Watch (live)
gh run watch 12345
gh run watch 12345 --exit-status

# Trigger
gh workflow run ci.yml
gh workflow run ci.yml --ref my-branch
gh workflow run ci.yml -f name=value -f other=value
echo '{"key":"val"}' | gh workflow run ci.yml --json

# Manage runs
gh run rerun 12345
gh run rerun 12345 --failed                  # rerun failed jobs only
gh run cancel 12345
gh run download 12345                        # download artifacts
gh run download 12345 --name artifact-name
gh run delete 12345

# Manage workflows
gh workflow list
gh workflow view ci.yml
gh workflow enable ci.yml
gh workflow disable ci.yml
```

### Search

```bash
# Issues
gh search issues "search terms"
gh search issues --assignee=@me --state=open
gh search issues --label bug --owner cli
gh search issues --comments=">100"
gh search issues -- -label:bug               # exclude label

# Pull requests
gh search prs --review-requested=@me --state=open
gh search prs --repo=cli/cli --draft
gh search prs --assignee=@me --merged
gh search prs -- -label:bug

# Code
gh search code "pattern"
gh search code "func main" --language=go
gh search code lint --filename package.json
gh search code panic --repo cli/cli
gh search code deque --language=python --owner=org

# Commits
gh search commits "fix bug" --author=user
gh search commits --repo=cli/cli --since=2024-01-01

# Repos
gh search repos "keyword" --language=python
gh search repos --owner=org --sort=stars
```

### Gists

```bash
gh gist create file.txt
gh gist create file.txt --public
gh gist create file.txt -d "Description"
gh gist list
gh gist view <id>
gh gist edit <id>
gh gist clone <id>
gh gist delete <id>
```

### GitHub API (Direct)

```bash
# REST API
gh api repos/{owner}/{repo}
gh api repos/{owner}/{repo}/issues
gh api repos/{owner}/{repo}/pulls/123/comments
gh api repos/{owner}/{repo}/releases --jq '.[0].tag_name'

# With parameters
gh api repos/{owner}/{repo}/issues/123/comments -f body="Comment"
gh api repos/{owner}/{repo}/issues -F title="Title" -F body="Body"

# HTTP methods
gh api repos/{owner}/{repo}/issues/123 -X PATCH -f state=closed
gh api repos/{owner}/{repo} -X DELETE

# Pagination
gh api repos/{owner}/{repo}/issues --paginate
gh api repos/{owner}/{repo}/issues --paginate --slurp --jq 'length'

# GraphQL
gh api graphql -f query='{ viewer { login } }'
gh api graphql -f query='query($owner:String!,$repo:String!) {
  repository(owner:$owner,name:$repo) { description }
}' -f owner=cli -f repo=cli

# Options
gh api endpoint --cache 3600s
gh api endpoint -i                           # include headers
gh api endpoint --verbose                    # full request/response
gh api endpoint --silent                     # suppress output
gh api endpoint -H "Accept: application/vnd.github+json"
```

### Configuration

```bash
gh config set editor vim
gh config set git_protocol ssh
gh config set prompt disabled
gh config get git_protocol
gh config list
gh config clear-cache
```

### Aliases

```bash
gh alias set co 'pr checkout'
gh alias set bugs 'issue list --label=bug'
gh alias list
gh alias delete co
```

## Global Flags

```bash
-R, --repo <[HOST/]OWNER/REPO>   # target a different repo
--json <fields>                   # output as JSON with specified fields
-q, --jq <expression>            # filter JSON output with jq
-t, --template <string>          # format JSON with Go template
-w, --web                        # open in browser
```

## JSON Output

Most list/view commands support `--json` for structured output:

```bash
# List available JSON fields
gh pr list --json

# Select specific fields
gh pr list --json number,title,state

# Filter with jq
gh pr list --json number,title --jq '.[] | select(.title | test("fix"))'

# Format with Go templates
gh pr list --json number,title -t '{{range .}}#{{.number}} {{.title}}{{"\n"}}{{end}}'
```

## Environment Variables

- `GITHUB_TOKEN` / `GH_TOKEN` - Authentication token
- `GH_REPO` - Default repository (OWNER/REPO)
- `GH_HOST` - Default GitHub host
- `GH_EDITOR` - Editor for text input
- `GH_BROWSER` - Browser for web commands
- `NO_COLOR` - Disable color output
- `GH_DEBUG` - Enable debug logging

## Examples

### Create a PR with full metadata

```bash
gh pr create \
  --title "Add login feature" \
  --body "Implements OAuth login flow" \
  --reviewer alice,bob \
  --label feature \
  --assignee @me \
  --milestone "v2.0"
```

### Review and merge a PR

```bash
gh pr checkout 123
gh pr checks --watch
gh pr review --approve -b "LGTM"
gh pr merge --squash --delete-branch
```

### Triage issues

```bash
gh issue list --assignee @me --state open --json number,title,labels
gh issue view 456
gh issue edit 456 --add-label "in-progress"
gh issue comment 456 --body "Working on this"
```

### Monitor CI

```bash
gh run list --workflow ci.yml --branch main --limit 5
gh run view 12345 --verbose
gh run view 12345 --log-failed
gh run rerun 12345 --failed
```

### Release workflow

```bash
gh release create v1.2.0 \
  --generate-notes \
  --title "v1.2.0" \
  ./dist/*.tar.gz
```
