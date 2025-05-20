# Human Presence Simulator 🏠

![Go](https://img.shields.io/badge/Go-1.21-blue)
![Test Coverage](https://img.shields.io/badge/coverage-dynamic-lightgrey?style=flat&logo=go)

A simple, extensible CLI utility to simulate human presence using IoT devices like Sonos (via gspotty), Hue, Tapo, and Google TV.

## ✨ Features

- 📦 Plugin-based architecture (`gspotty`, `hue`, `tapo`)
- 🗓 Declarative routine config (YAML or JSON)
- 🧪 Strong unit test support with mocks
- 🔧 Structured telemetry via JSON logs
- 🧰 CLI flags for validation, introspection

## 🚀 Getting Started

```bash
git clone https://github.com/iamgaru/human-presence-sim.git
cd human-presence-sim
make build
./bin/simulator --list-plugins
```

## 🧪 Running Tests

```bash
make test         # run unit tests
make race         # run tests with race detector
make coverage     # generate coverage report
```

## 📂 Routine Example

```yaml
name: "Evening Routine"
steps:
  - device: "hue"
    action: "set_light"
    args:
      room: "living"
      color: "warm"
      brightness: 75
  - delay: "5s"
    device: "gspotty"
    action: "play_playlist"
    args:
      name: "Dinner Chill"
```

## 📄 License

MIT © iamgaru

## 🧾 JSON Routine Example

```json
{
  "name": "Evening Routine",
  "steps": [
    {
      "device": "hue",
      "action": "set_light",
      "args": {
        "room": "living",
        "color": "warm",
        "brightness": 75
      }
    },
    {
      "delay": "5s",
      "device": "gspotty",
      "action": "play_playlist",
      "args": {
        "name": "Dinner Chill"
      }
    }
  ]
}
```
