# fm

How many times did you just want to open your current directory from the terminal in your default file manager? `fm` does it for you.

Works on Linux, macOS, and Windows — no configuration required.

---

## Installation

You need [Go](https://go.dev) installed (1.21 or newer).

```bash
git clone https://github.com/tissanr/fm.git
cd fm
make install
```

This builds the binary and installs it to `~/.local/bin/fm` — no `sudo` required.

To install to a different location pass `PREFIX`:

```bash
make install PREFIX=/usr/local
```

Other make targets:

| Command | What it does |
|---------|--------------|
| `make build` | Compile the binary into the current directory |
| `make install` | Build and install to `$PREFIX/bin` (default: `~/.local/bin`) |
| `make uninstall` | Remove the installed binary |
| `make clean` | Remove the local build artifact |

> Make sure `~/.local/bin` is on your `PATH`. Most Linux distros include it by default. If `fm` is not found after install, add this to your `~/.bashrc` or `~/.zshrc`:
> ```bash
> export PATH="$HOME/.local/bin:$PATH"
> ```

> **Windows:** Use the included PowerShell script instead of `make`:
>
> ```powershell
> git clone https://github.com/tissanr/fm.git
> cd fm
> .\install.ps1
> ```
>
> This builds `fm.exe` and installs it to `%LOCALAPPDATA%\Microsoft\WindowsApps`, which is already on your PATH without needing admin rights.
>
> Other commands:
>
> | Command | What it does |
> |---------|--------------|
> | `.\install.ps1` | Build and install (default) |
> | `.\install.ps1 build` | Compile only |
> | `.\install.ps1 uninstall` | Remove the installed binary |
> | `.\install.ps1 clean` | Remove the local build artifact |
>
> To install to a custom folder: `.\install.ps1 -Target "C:\my\folder"`
>
> If PowerShell blocks the script, run this once first:
> `Set-ExecutionPolicy -Scope CurrentUser RemoteSigned`

---

## Usage

```
fm [path]
```

| Command | What it does |
|---------|--------------|
| `fm` | Opens the current directory |
| `fm .` | Same as above |
| `fm /some/folder` | Opens `/some/folder` |
| `fm ..` | Opens the parent directory |
| `fm --help` | Shows usage and the config file location |

---

## Configuration

You can override the file manager that `fm` uses by creating a config file.
`fm --help` prints the exact path for your system.

| OS | Config file location |
|----|----------------------|
| Linux | `~/.config/fm/config.toml` (or `$XDG_CONFIG_HOME/fm/config.toml`) |
| macOS | `~/Library/Application Support/fm/config.toml` |
| Windows | `%APPDATA%\fm\config.toml` |

### Config file format

```toml
file_manager = "nautilus"
```

Replace `"nautilus"` with any file manager binary that is on your `PATH`,
or use an absolute path.

### Examples

```toml
# Linux — use Thunar instead of the xdg-open default
file_manager = "thunar"

# macOS — use Finder explicitly (this is also the default)
file_manager = "open"

# Windows — use Total Commander
file_manager = "C:\\Program Files\\totalcmd\\TOTALCMD64.EXE"
```

If the config file does not exist, `fm` falls back to the OS default:

| OS | Default command |
|----|-----------------|
| Linux | `xdg-open` |
| macOS | `open` |
| Windows | `explorer` |

---

## Error messages

| Message | Cause |
|---------|-------|
| `fm: path does not exist: …` | The given path does not exist on disk |
| `fm: not a directory: …` | The path exists but is a file, not a folder |
| `fm: could not launch "…": …` | The file manager binary was not found or failed to start |
| `fm: warning: …` | Config file exists but could not be read |
