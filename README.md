# RazerBatteryTray

Linux system tray application that displays the battery level of your Razer devices, powered by the [OpenRazer](https://openrazer.github.io/) daemon.
Go variant of [RazerBatteryTray](https://github.com/HoroTW/RazerBatteryTray).

![Platform](https://img.shields.io/badge/platform-linux-lightgrey) ![Go](https://img.shields.io/badge/go-1.22%2B-00ADD8)

## Features

- Tray icon with an always-visible battery percentage
- Color-coded levels:
  - Green — normal (> 20%)
  - Amber — low (≤ 20%)
  - Red — critical (≤ 10%)
  - Teal — charging
- Device picker menu — switch between multiple Razer devices
- Automatic fallback to the device list if the selected device stops responding

## Requirements

- Linux with a system tray host (e.g. GNOME with AppIndicator support, KDE Plasma)
- [OpenRazer](https://openrazer.github.io/) correctly installed
- Go 1.22+ (only needed to build)

Verify the daemon is reachable:

```bash
busctl --user list | grep org.razer
```


## Build & Run

```bash
go build -o razerbatterytray .
./razerbatterytray
```

The binary targets Linux by default; cross-compile explicitly with:

```bash
GOOS=linux GOARCH=amd64 go build -o razerbatterytray .
```

```bash
chmod +x razerbatterytray
./razerbatterytray
```


### Autostart with systemd

To start the tray automatically at login, create a **user** service at `~/.config/systemd/user/razerbatterytray.service`:

```ini
[Unit]
Description=Razer Battery Tray
PartOf=graphical-session.target
After=graphical-session.target

[Service]
ExecStart=<ExecFilePath>
Restart=on-failure
RestartSec=5

[Install]
WantedBy=graphical-session.target
```

Adjust `ExecStart` to wherever you placed the binary.
Then enable it:

```bash
systemctl --user daemon-reload
systemctl --user enable --now razerbatterytray.service
```

Check status and logs:

```bash
systemctl --user status razerbatterytray.service
journalctl --user -u razerbatterytray -f
```

## License And Stuff
RAZER is a trademark of Razer Inc. This application is an independent, unofficial project and is not affiliated with or endorsed by Razer Inc.
