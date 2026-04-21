package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func configPath() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "fm", "config.toml")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "fm", "config.toml")
	default:
		base := os.Getenv("XDG_CONFIG_HOME")
		if base == "" {
			base = filepath.Join(os.Getenv("HOME"), ".config")
		}
		return filepath.Join(base, "fm", "config.toml")
	}
}

// loadFileManager reads file_manager from the config file.
// Returns "" if the config doesn't exist or the key is not set.
func loadFileManager() string {
	path := configPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
		if key == "file_manager" && value != "" {
			return value
		}
	}
	return ""
}

func defaultLauncher() string {
	switch runtime.GOOS {
	case "windows":
		return "explorer"
	case "darwin":
		return "open"
	default:
		return "xdg-open"
	}
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println("Usage: fm [path]")
		fmt.Println()
		fmt.Println("  Opens a directory in the default file manager.")
		fmt.Println("  Defaults to the current directory when no path is given.")
		fmt.Println()
		fmt.Printf("  Config file: %s\n", configPath())
		os.Exit(0)
	}

	target := "."
	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	abs, err := filepath.Abs(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fm: could not resolve path: %v\n", err)
		os.Exit(1)
	}

	info, err := os.Stat(abs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fm: path does not exist: %s\n", abs)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "fm: not a directory: %s\n", abs)
		os.Exit(1)
	}

	launcher := loadFileManager()
	if launcher == "" {
		launcher = defaultLauncher()
	}

	cmd := exec.Command(launcher, abs)
	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "fm: could not launch %q: %v\n", launcher, err)
		os.Exit(1)
	}
}
