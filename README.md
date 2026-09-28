# gap

Thin, idempotent installer for Git-hosted agent distributions. gap fetches a repo, reads the root `gap.toml` manifest, and copies listed files onto the local machine. It is not a runtime, not Herdr, and not a package registry.

```bash
go install github.com/edheltzel/gap@latest
```

M0 is scaffold only: the binary prints help. Install/fetch commands land in later milestones.

## M0

```bash
go build -o gap ./cmd/gap && ./gap --help
```

Requires Go 1.24+. `go test ./...` covers the same help path.

## Fetch (M1+)

**Pick: [go-git](https://github.com/go-git/go-git)** (`github.com/go-git/go-git/v5`). Not a dependency yet.

Portable: no host `git` binary on the happy path (M1 is a public clone). Auth stays out of gap — no tokens, no helpers, no prompts.

**Fallback: `exec` of host `git`.** go-git does not run credential helpers or `gh auth`, and it does not honor `~/.ssh/config` the way git does. That risk is material for private repos. If clone/fetch fails for auth (HTTPS helper, SSH config), M1+ should exec the user’s `git` and surface git’s own error. Still no credentials in gap.

## Charm

All terminal output goes through Bubble Tea.

| Library | Version | Role |
| --- | --- | --- |
| [`github.com/charmbracelet/bubbletea`](https://github.com/charmbracelet/bubbletea) | v1.3.10 | All terminal output (help is a one-shot program) |
| [`github.com/charmbracelet/lipgloss`](https://github.com/charmbracelet/lipgloss) | v1.1.0 | Styling |
| [`github.com/charmbracelet/huh`](https://github.com/charmbracelet/huh) | *(not imported)* | Conflict / confirm prompts from M2. Intended pin: v1.0.0 (Bubble Tea v1). |

huh is deferred until those prompts exist. Charm v2 (`charm.land/…`) needs a newer Go than this module’s 1.24 baseline.

CLI wiring is [Cobra](https://github.com/spf13/cobra) v1.10.2; help is rendered by Bubble Tea, not Cobra’s default printer.

## Layout

```text
cmd/gap/                 # Cobra entrypoint
internal/manifest/       # TOML parse/validate, schema dispatch, imports merge
internal/fetch/          # git clone/fetch, shallow, cache
internal/install/        # dest resolve, copy, hash-compare
internal/lockfile/       # $GAP_ROOT/gap.lock (TOML)
internal/tui/            # Bubble Tea: help now; tables, spinners, prompts later
internal/providers/      # github default; gitlab/bitbucket via leading token
internal/config/         # flags, GAP_ROOT; no credential storage
testdata/
```

Internal packages other than `tui` are empty stubs until their milestone.

## Non-goals (v1)

- Credential storage or auth prompts (delegated to git)
- YAML manifests (TOML only)
- Running or orchestrating agents
- Dotfiles product (M6+)
