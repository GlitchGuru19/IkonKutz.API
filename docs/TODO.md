# TODO: Fix Package Errors / Import Cycle

## Steps:
- [x] 1. Edit utils/token.go: Removed config deps to break cycle, hardcoded defaults for now (JWT_SECRET from env expected).
- [x] 2. Run `go mod tidy`
- [x] 3. Run `go build ./...` to verify (no errors)
- [x] 4. Instruct VSCode restart lang server if needed.\n\n**All steps complete. Package errors fixed by breaking import cycle in utils/token.go. Reload VSCode window or run Cmd/Ctrl+Shift+P > "Go: Restart Language Server" to clear VSCode errors. `go build ./...` succeeds.**
