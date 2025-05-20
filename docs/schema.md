# 📘 Routine Configuration Schema

This defines the schema used to configure your IoT routine simulations.

## Top-Level Fields

| Field | Type   | Description                      |
|-------|--------|----------------------------------|
| name  | string | Name of the routine              |
| at    | string | Optional time (HH:MM) to start   |
| steps | array  | Sequence of device steps to run  |

## Step Fields

| Field   | Type                | Description                                  |
|---------|---------------------|----------------------------------------------|
| device  | string              | Device key (e.g., `gspotty`, `hue`, `tapo`) |
| action  | string              | Action name supported by the plugin          |
| args    | map<string, any>    | Key-value args for that action               |
| delay   | string (optional)   | Wait time before this step (e.g. "5s")     |

## Example (YAML)

```yaml
name: Morning Routine
at: "07:30"
steps:
  - device: hue
    action: set_light
    args:
      room: bedroom
      brightness: 60
      color: warm
  - delay: "10s"
    device: gspotty
    action: play_playlist
    args:
      name: Wake Up Mix
```
