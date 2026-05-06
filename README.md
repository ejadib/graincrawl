# graincrawl

`graincrawl` is a local-first archive tool for Granola notes, transcripts,
summaries, panels, and meeting metadata.

It stores a private SQLite archive under `~/.config/graincrawl`, exposes stable
JSON command output for automation, and keeps Granola's private surfaces behind
explicit source adapters.

## status

Early implementation. The default target is read-only sync through Granola's
desktop private API token, with plaintext desktop cache as an offline fallback.

## core commands

```bash
graincrawl doctor
graincrawl sync --source private-api
graincrawl sync --source desktop-cache
graincrawl notes --json
graincrawl note get <id> --json
graincrawl transcripts get <id> --json
graincrawl panels get <id> --json
graincrawl export markdown --out ./granola-notes
```

## source policy

`graincrawl` never writes to Granola's local profile. It reads from copied local
files or Granola read endpoints only.

Supported source names:

- `private-api`
- `desktop-cache`
- `public-api`
- `companion-cli`
- `encrypted-json`
- `opfs`

The encrypted and OPFS sources require explicit unlock commands before they can
touch macOS Keychain, Electron safeStorage, IndexedDB, or OPFS state.

## development

```bash
go test ./...
go vet ./...
go run ./cmd/graincrawl --help
```
