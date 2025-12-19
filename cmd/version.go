package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	version        string
	buildTime      string
	versionOutputJSON bool
	versionOutputYAML bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		v := version
		bt := buildTime

		// If version is empty, try to get it from build info
		if v == "" {
			if info, ok := debug.ReadBuildInfo(); ok {
				if info.Main.Version != "" && info.Main.Version != "(devel)" {
					v = info.Main.Version
				} else {
					// Try to get version from VCS info
					for _, setting := range info.Settings {
						if setting.Key == "vcs.revision" {
							if len(setting.Value) > 7 {
								v = setting.Value[:7]
							} else {
								v = setting.Value
							}
							break
						}
					}
					if v == "" {
						v = "dev"
					}
				}
			} else {
				v = "unknown"
			}
		}

		// If buildTime is empty, try to get it from build info
		if bt == "" {
			if info, ok := debug.ReadBuildInfo(); ok {
				for _, setting := range info.Settings {
					if setting.Key == "vcs.time" {
						bt = setting.Value
						break
					}
				}
			}
		}

		output := map[string]string{
			"version":   v,
			"build_time": bt,
		}

		if versionOutputJSON {
			encoder := json.NewEncoder(os.Stdout)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(output); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to output JSON: %v\n", err)
				os.Exit(1)
			}
			return
		}

		if versionOutputYAML {
			encoder := yaml.NewEncoder(os.Stdout)
			defer encoder.Close()
			if err := encoder.Encode(output); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to output YAML: %v\n", err)
				os.Exit(1)
			}
			return
		}

		fmt.Printf("sfvbrr version: %s\n", v)
		if bt != "" && bt != "unknown" {
			fmt.Printf("Build Time:    %s\n", bt)
		}
	},
	DisableFlagsInUseLine: true,
}

func SetVersion(v, bt string) {
	if v == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok {
			v = info.Main.Version
		}
	}
	version = v
	buildTime = bt
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVar(&versionOutputJSON, "json", false, "Output version information in JSON format")
	versionCmd.Flags().BoolVar(&versionOutputYAML, "yaml", false, "Output version information in YAML format")
	versionCmd.MarkFlagsMutuallyExclusive("json", "yaml")
	versionCmd.SetUsageTemplate(`Usage:
  {{.CommandPath}}

Prints the version and build time information for sfvbrr.
`)
}
