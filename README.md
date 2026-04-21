# fm

A tiny command-line tool that opens a folder in your default file manager.

Works on Linux, macOS, and Windows — no configuration required.

---

## Installation

You need [Go](https://go.dev) installed (1.21 or newer).

```bash
cd /path/to/fm
go build -o fm .
```

Then move the binary somewhere on your `PATH`:

```bash
# Linux / macOS
mv fm /usr/local/bin/fm

# Windows — move fm.exe to any folder that is in your PATH
```

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
