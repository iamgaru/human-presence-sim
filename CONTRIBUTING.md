# 🤝 Contributing to Human Presence Simulator

Thank you for considering a contribution!

## 🧱 Adding a New Plugin

1. Create a new folder in `plugins/` (e.g. `plugins/mydevice`)
2. Implement the `types.Plugin` interface:
    ```go
    func (p *MyPlugin) Name() string { return "mydevice" }
    func (p *MyPlugin) Init(cfg map[string]interface{}) error { ... }
    func (p *MyPlugin) PerformAction(action string, args map[string]interface{}) error { ... }
    ```
3. Register the plugin in `internal/registry/registry.go`
4. Add tests in `plugins/mydevice/mydevice_test.go`
5. Update documentation and `README.md` if needed

## 🔧 Setup Instructions

```sh
make install       # builds and installs the CLI to /usr/local/bin
make test          # run tests
make coverage      # check test coverage
```

## 📦 Pull Requests

- Keep PRs focused
- Add tests and documentation
- Use semantic commit messages (feat:, fix:, chore:, etc.)

