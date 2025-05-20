# 📦 human-presence-sim v0.1.0

Initial release of the Human Presence Simulator CLI.

### ✨ Features

- Plugin-based architecture: `gspotty`, `hue`, `tapo`
- Declarative YAML/JSON routines for simulating human presence
- Structured JSON logging via `telemetry` module
- CLI tools: `--list-plugins`, `--validate-config`
- Full test coverage with mocks and utility helpers
- CI-ready with GitHub Actions
- `.env` support for plugin auth/secrets

### 🛠 Usage

```sh
make build
./bin/simulator --list-plugins
./bin/simulator config/example.yaml
```

### 🧪 Test

```sh
make test       # unit tests
make coverage   # view coverage.html
```

