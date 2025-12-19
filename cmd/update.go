package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blang/semver/v4"
	"github.com/creativeprojects/go-selfupdate"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	updateOutputJSON bool
	updateOutputYAML bool
)

var updateCmd = &cobra.Command{
	Use:                   "update",
	Short:                 "Update sfvbrr",
	Long:                  `Update sfvbrr to latest version.`,
	RunE:                  runUpdate,
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVar(&updateOutputJSON, "json", false, "Output update information in JSON format")
	updateCmd.Flags().BoolVar(&updateOutputYAML, "yaml", false, "Output update information in YAML format")
	updateCmd.MarkFlagsMutuallyExclusive("json", "yaml")
	updateCmd.SetUsageTemplate(`Usage:
  {{.CommandPath}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
`)
}

func runUpdate(cmd *cobra.Command, args []string) error {
	currentVersion, err := semver.ParseTolerant(version)
	if err != nil {
		return fmt.Errorf("could not parse version: %w", err)
	}

	latest, found, err := selfupdate.DetectLatest(cmd.Context(), selfupdate.ParseSlug("autobrr/sfvbrr"))
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %w", err)
	}
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from github repository", "autobrr/sfvbrr", version)
	}

	latestVersion, err := semver.ParseTolerant(latest.Version())
	if err != nil {
		return fmt.Errorf("could not parse latest version: %w", err)
	}

	output := map[string]interface{}{
		"current_version": version,
		"latest_version":  latest.Version(),
		"updated":         false,
	}

	if latestVersion.LTE(currentVersion) {
		if updateOutputJSON || updateOutputYAML {
			output["message"] = fmt.Sprintf("Current binary is the latest version: %s", version)
		} else {
			fmt.Printf("Current binary is the latest version: %s\n", version)
		}
	} else {
		exe, err := selfupdate.ExecutablePath()
		if err != nil {
			return fmt.Errorf("could not locate executable path: %w", err)
		}

		if err := selfupdate.UpdateTo(cmd.Context(), latest.AssetURL, latest.AssetName, exe); err != nil {
			return fmt.Errorf("error occurred while updating binary: %w", err)
		}

		output["updated"] = true
		if updateOutputJSON || updateOutputYAML {
			output["message"] = fmt.Sprintf("Successfully updated to version: %s", latest.Version())
		} else {
			fmt.Printf("Successfully updated to version: %s\n", latest.Version())
		}
	}

	if updateOutputJSON {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(output)
	}

	if updateOutputYAML {
		encoder := yaml.NewEncoder(os.Stdout)
		defer encoder.Close()
		return encoder.Encode(output)
	}

	return nil
}
