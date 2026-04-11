# Hot Reload with CompileDaemon in development

## Installation
First, install CompileDaemon:

go install github.com/githubnemo/CompileDaemon@latest


## Usage
In the project root, run:


CompileDaemon -build="go build -o main ." -command="./main"


## How it works
- **-build="go build -o main ."**: Builds the `main` binary from the current module (IkonKutz.API).
- **-command="./main"**: Runs the built binary.
- Watches for Go file changes, rebuilds automatically, and restarts the server for hot reload.

Make code changes, save, and see instant reload!

## Tips
- For custom build flags, adjust the `-build` flag (e.g., `-build="go build -tags=dev -o main ."`).
- Stop with Ctrl+C.

