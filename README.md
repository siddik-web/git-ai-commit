# Git AI Commit - Intelligent Commit Automation

Git AI Commit is a powerful CLI tool that leverages local AI models (via Ollama) to automatically generate meaningful commit messages and push your changes with a single command. Perfect for developers who want to maintain clean commit history without the manual effort.

## Key Features

- ✨ **AI-Powered Commit Messages** - Context-aware commit messages generated from your code changes
- 🔒 **100% Local Processing** - Your code never leaves your machine
- ⚡️ **Single-Command Workflow** - Stage, commit, and push with one command
- 🤖 **Multi-Model Support** - Works with any Ollama-compatible model (Mistral, Llama 3, CodeLlama, etc.)
- 📝 **Conventional Commits** - Follows industry-standard commit conventions
- 🛡️ **Secure** - Uses your existing Git credentials

## Installation

### Prerequisites
1. [Install Go](https://go.dev/dl/) (v1.16+)
2. [Install Ollama](https://ollama.com/)
3. Download a language model:
   ```bash
   ollama pull mistral  # Recommended model
   ```

### Install Git AI Commit
```bash
# Clone the repository
git clone https://github.com/yourusername/git-ai-commit.git
cd git-ai-commit

# Build and install
go build -o git-ai-commit
sudo mv git-ai-commit /usr/local/bin/
```

## Usage

```bash
# Stage your changes
git add .

# Generate commit message and push
git-ai-commit

# Use a specific model
git-ai-commit --model llama3

# See help
git-ai-commit --help
```

## Example Workflow

```bash
$ git add .
$ git-ai-commit

Generated commit message:

feat: implement auto-commit functionality using Ollama API

Successfully committed and pushed changes!
```

## Configuration

Customize the behavior by modifying these elements in the code:

1. **Prompt Engineering** - Modify the prompt template in `generateCommitMessage()`:
   ```go
   prompt := fmt.Sprintf(`Generate a concise git commit message...`)
   ```

2. **Commit Message Cleaning** - Adjust the cleaning logic in `cleanCommitMessage()`:
   ```go
   func cleanCommitMessage(msg string) string {
       // Your custom cleaning logic
   }
   ```

3. **Commit Rules** - Add validation checks before committing

## Supported Models

Any Ollama-compatible model can be used:
- `mistral` (default)
- `llama3`
- `codellama`
- `phi3`
- `gemma`

Find more models at [Ollama Library](https://ollama.com/library)

## FAQ

### Q: Is my code sent to external servers?
A: No! All processing happens locally using your Ollama installation. Your code never leaves your machine.

### Q: What if I want to review the commit before pushing?
A: Currently the tool commits and pushes automatically. You can modify the code to remove the push step if you prefer manual pushing.

### Q: Can I use this with private repositories?
A: Absolutely! The tool uses your existing Git credentials and works with any repository you have access to.

### Q: How do I uninstall?
```bash
sudo rm /usr/local/bin/git-ai-commit
```

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a pull request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

**Git AI Commit** - Never write a commit message again! ✨🤖💾
