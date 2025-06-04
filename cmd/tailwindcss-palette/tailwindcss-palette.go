package main

import (
	"os"
	"runtime/debug"

	"github.com/claytonchew/tailwindcss-palette-go/internal/clicmd"
	"github.com/claytonchew/tailwindcss-palette-go/internal/version"
)

func main() {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				if version.CommitHash == "unknown" {
					version.CommitHash = setting.Value
				}
			case "vcs.time":
				if version.BuildDate == "unknown" {
					version.BuildDate = setting.Value
				}
			}
		}
	}

	code := clicmd.Main()
	os.Exit(int(code))
}
