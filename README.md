# gogent-cli

> A terminal-native coding agent built in Go.

gogent-cli is a CLI tool that understands your codebase — scanning repositories, reading files, searching code, and executing tasks through a structured agent loop. Built with zero external dependencies.

---

## Handwritten Codebase

Every line of code in this repository is handwritten. AI was used exclusively for research and technical guidance — no AI-generated code exists in this codebase.

---

## Features

| Command | Description |
|---|---|
| `scan` | Recursively walks the repository and prints a file tree, honoring `.gitignore` rules and sensible defaults |
| `read <path>` | Reads a file with line numbers. Enforces a 100KB size cap |

---

## Installation

**Prerequisites:** Go 1.23+

```bash
git clone https://github.com/<your-username>/gogent-cli.git
cd gogent-cli
go build -o gogent .
```

---

## Usage

```bash
# Print the repository file tree
./gogent scan

# Read a file with line numbers
./gogent read path/to/file.go
```

---

## Project Structure

```
gogent-cli/
├── main.go              # CLI routing and entrypoint
├── scan.go              # Recursive repository scanner
├── scan_test.go         # Scanner tests
├── gitignore.go         # .gitignore parser
├── read_file.go         # File reader with line numbers and size guard
└── read_file_test.go    # File reader tests
```

---

## Roadmap

- [x] Repository scanner with `.gitignore` support
- [x] `read_file` — line-numbered output with size guard
- [ ] `search` — grep-style pattern search across the repository
- [ ] Safe command runner with explicit approval flow
- [ ] Patch proposal and diff display
- [ ] Patch approval and apply flow
- [ ] Fake scripted model adapter
- [ ] Real LLM adapter — Anthropic, OpenAI, Gemini
- [ ] Session logging
- [ ] Context window management

---

## Development

```bash
# Run tests
go test ./...

# Format code
gofmt -w .
```

---

## License

MIT
